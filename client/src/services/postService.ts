import type { Post, PostComment, Tag } from '@/interfaces/post'
import pbClient from '@/pocketbase'

export async function getAllPostsPaginated(
  page: number,
  limit: number,
): Promise<{
  totalItems: number
  totalPages: number
  posts: Post[]
}> {
  const postDetails = await pbClient
    .collection('posts_author_vw')
    .getList(page, limit, { expand: 'tags', sort: '-created' })

  return {
    totalItems: postDetails.totalItems,
    totalPages: postDetails.totalPages,
    posts: postDetails.items.map(parsePost),
  }
}

export async function getPostByUserIdPaginated(
  userId: string,
  page: number,
  limit: number,
): Promise<{
  totalItems: number
  totalPages: number
  posts: Post[]
}> {
  const postDetails = await pbClient
    .collection('posts_author_vw')
    .getList(page, limit, {
      expand: 'tags',
      sort: '-created',
      filter: pbClient.filter('userId = {:userId}', { userId }),
    })

  return {
    totalItems: postDetails.totalItems,
    totalPages: postDetails.totalPages,
    posts: postDetails.items.map(parsePost),
  }
}

export async function getPostById(postId: string): Promise<Post> {
  const postDetails = await pbClient
    .collection('posts_author_vw')
    .getOne(postId, { expand: 'tags' })

  return parsePost(postDetails)
}

/**
 *
 * @param title
 * @param content
 * @param tags
 * @returns userId of the newly created post
 */
export async function createNewPost(
  title: string,
  content: string,
  tags: string[],
): Promise<string> {
  const resultPost = await pbClient.collection('posts').create({
    userId: pbClient.authStore?.model?.id,
    title,
    content,
    tags,
  })

  return resultPost.id
}

export async function updatePost(
  postId: string,
  fields: { title: string; content: string; tagIds: string[] },
): Promise<Post> {
  await pbClient.collection('posts').update(postId, fields)

  const postDetails = await pbClient
    .collection('posts_author_vw')
    .getOne(postId)

  return parsePost(postDetails)
}

function parsePost(data: unknown): Post {
  const postData = data as {
    id: string
    userId: string
    title: string
    content: string
    firstName: string
    lastName: string
    expand?: {
      tags: {
        id: string
        name: string
      }[]
    }
    created: string
    updated: string
  }

  const post: Post = {
    id: postData.id,
    userId: postData.userId,
    title: postData.title,
    content: postData.content,
    firstName: postData.firstName,
    lastName: postData.lastName,
    tags: [],
    created: new Date(postData.created),
    updated: new Date(postData.updated),
  }

  if (postData.expand) {
    post.tags = postData.expand.tags.map(tag => ({
      id: tag.id,
      name: tag.name,
    }))
  }

  return post
}

export async function deletePost(postId: string): Promise<void> {
  await pbClient.collection('posts').delete(postId)
}

////////////////////////////////////
// Post Comments Functions
////////////////////////////////////
export async function createNewComment(
  postId: string,
  comment: string,
): Promise<PostComment> {
  const resultComment = await pbClient.collection('comments').create({
    userId: pbClient.authStore?.model?.id,
    postId: postId,
    content: comment,
  })

  const commentDetails = await pbClient
    .collection('comments_author_vw')
    .getOne(resultComment.id)

  return parsePostComment(commentDetails)
}

export async function getPostCommentsPaginated(
  postId: string,
  page: number,
  limit: number,
): Promise<{
  totalItems: number
  totalPages: number
  comments: PostComment[]
}> {
  const response = await pbClient
    .collection('comments_author_vw')
    .getList(page, limit, {
      filter: `postId='${postId}'`,
      sort: '-created',
    })

  return {
    totalItems: response.totalItems,
    totalPages: response.totalPages,
    comments: response.items.map(parsePostComment),
  }
}

function parsePostComment(data: unknown): PostComment {
  const postData = data as {
    id: string
    content: string
    firstName: string
    lastName: string
    created: string
    updated: string
  }

  return {
    id: postData.id,
    content: postData.content,
    firstName: postData.firstName,
    lastName: postData.lastName,
    created: new Date(postData.created),
    updated: new Date(postData.updated),
  }
}

////////////////////////////////////
// Tags Functions
////////////////////////////////////
export async function getPostTagsPage(
  page: number,
  limit: number,
): Promise<{
  totalItems: number
  totalPages: number
  tags: Tag[]
}> {
  const response = await pbClient.collection('tags').getList(page, limit)

  return {
    totalItems: response.totalItems,
    totalPages: response.totalPages,
    tags: response.items.map(parseTag),
  }
}

export async function searchPostTags(query: string): Promise<Tag[]> {
  const trimmedQuery = query.trim()
  if (trimmedQuery !== '') {
    {
      const response = await pbClient.collection('tags').getList(1, 10, {
        filter: pbClient.filter('name ~ {:name}', { name: trimmedQuery }),
      })

      return response.items.map(parseTag)
    }
  } else {
    return []
  }
}

export async function createNewTag(tagName: string): Promise<Tag> {
  const resultTag = await pbClient.collection('tags').create({ name: tagName })

  return parseTag(resultTag)
}

function parseTag(data: unknown): Tag {
  const tagData = data as {
    id: string
    name: string
  }

  return {
    id: tagData.id,
    name: tagData.name,
  }
}
