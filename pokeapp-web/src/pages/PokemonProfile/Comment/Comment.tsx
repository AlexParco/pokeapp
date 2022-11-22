import { useAuth } from "@/context/auth.context"
import { addCommentService, getCommentsByPokeId } from "@/service/api"
import { Comment as CommentType } from "@/types/comment.type"
import { Button, FormControl, Heading, Input } from "@chakra-ui/react"
import { useEffect, useState } from "react"
import { CommentList } from "./CommentList"

const Comment = ({ pokeId }: { pokeId: string }) => {
  const [comments, setComments] = useState<CommentType[]>([])
  const [body, setBody] = useState<string>("")
  const { logout } = useAuth()
  const { token } = useAuth()

  useEffect(() => {
    getCommentsByPokeId(pokeId)
      .then(comments => {
        setComments(comments)
      })
      .catch(_ => {
        logout()
      })
  }, [setComments])

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()

    addCommentService({ pokemon_id: Number(pokeId), body }, token)
      .then(comment => {
        if (!comments) {
          setComments([comment])
        } else {
          setComments([comment, ...comments])
        }
      })
      .catch(err => {
        console.log(err)
      })

    setBody("")
  }

  return (
    <>
      <Heading as='h2' size='lg' mb={6}>Comments</Heading>
      <form
        onSubmit={handleSubmit}
      >
        <FormControl
          mb={6}
          display='flex'
          gap={2}
        >
          <Input
            value={body}
            onChange={(e) => setBody(e.target.value)}
          />
          <Button type="submit" >Enviar</Button>
        </FormControl>
      </form>
      {
        !comments ? <h1>no comments</h1>
          :
          <CommentList
            comments={comments}
          />
      }
    </>
  )
}

export default Comment