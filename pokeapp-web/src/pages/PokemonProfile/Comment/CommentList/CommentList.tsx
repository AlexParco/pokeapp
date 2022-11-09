import { Comment } from "@/types/comment.type"
import { Box, Flex, Tag, Text } from "@chakra-ui/react"

type Props = {
  comments: Comment[]
}

const CommentList = ({ comments }: Props) => {
  console.log(comments)
  return (
    <Box >
      {comments.map(comment => (
        <Flex key={comment.comment_id}
          borderRadius="10"
          border='1px'
          borderColor='blackAlpha.300'
          p='2'
          mb={3}
          align='center'
          gap='5px'
        >
          <Text as='b'>
            {comment.username}:
          </Text>
          <Text>
            {comment.body}
          </Text>
        </Flex>
      ))
      }
    </Box >
  )
}

export default CommentList 