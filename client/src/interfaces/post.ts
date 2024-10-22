export interface Post {
  id: string
  title: string
  content: string
  firstName: string
  lastName: string
  tags: Tag[]
  created: string
  updated: string
}

export interface Tag {
  id: string
  name: string
}
