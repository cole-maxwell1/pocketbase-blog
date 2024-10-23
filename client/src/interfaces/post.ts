export interface Post {
  id: string
  userId: string
  title: string
  content: string
  firstName: string
  lastName: string
  tags: Tag[]
  created: Date
  updated: Date
}

export interface Tag {
  id: string
  name: string
}

export interface PostComment {
  id: string
  content: string
  firstName: string
  lastName: string
  created: Date
  updated: Date
}
