import type { Author } from '@/interfaces/author'
import pbClient from '@/pocketbase'

export async function getAuthorById(authorId: string): Promise<Author> {
  const authorDetails = await pbClient.collection('authors_vw').getOne(authorId)
  return parseAuthor(authorDetails)
}

export async function updateAuthorBio(editedBio: string): Promise<void> {
  await pbClient
    .collection('users')
    .update(pbClient.authStore?.model?.id, { bio: editedBio })
}

function parseAuthor(data: unknown): Author {
  const authorDetails = data as {
    id: string
    firstName: string
    lastName: string
    bio: string
    created: string
    updated: string
  }

  return {
    id: authorDetails.id,
    firstName: authorDetails.firstName,
    lastName: authorDetails.lastName,
    bio: authorDetails.bio,
    created: new Date(authorDetails.created),
    updated: new Date(authorDetails.updated),
  }
}
