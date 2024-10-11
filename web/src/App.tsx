import { useEffect, useState, FC } from 'react'
import { Box, Stack, SxProps, Typography, Container, Paper, IconButton, TextField, Button, Backdrop } from '@mui/material'
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import CloseIcon from '@mui/icons-material/Close';

import axios from 'axios'

interface FormCommandProps {
  handleCreateBook?: (event: React.FormEvent<HTMLFormElement>) => Promise<void>
  setData: (key: string, value: string) => void
  book: Book
}

const FormCommand: FC<FormCommandProps> = ({
  handleCreateBook,
  setData,
  book
}) => {
  const CommandContainerSx: SxProps = {display:"flex", justifyContent:"center", bgcolor: "#ffff", width: 500, marginTop: 5}
  return (
    <Container component={'form'} sx={CommandContainerSx} onSubmit={handleCreateBook}>
      <Stack direction={"row"} spacing={2}>
        <TextField id="outlined-basic" label="Title" variant="outlined" size='small' type='text' value={book.title} onChange={(e)=>{setData('title', e.target.value)}}/>
        <TextField id="outlined-basic" label="Author" variant="outlined" size='small' type='text' value={book.author} onChange={(e)=>{setData('author', e.target.value)}}/>
        <Button variant="contained" type='submit'>บันทึก</Button>
      </Stack>
    </Container>
  )
}

const hostname = "localhost"

interface Book {
  id: number
  title: string
  author: string
  [key: string]: string | number
}

function getBooks() {
  return axios.request({
    method: 'get',
    url: `http://${hostname}:8080/books`,
  })
}

async function deleteBook(id: number) {
  return axios.request({
    method: 'delete',
    url: `http://${hostname}:8080/book?id=${id}`,
  })
}

function App() {
  const [books, setBooks] = useState<Book[]>([])
  const [book, setBook] = useState<Book>({id: 0, title: "", author: ""})
  const [editBook, setEditBook] = useState<Book>({id: 0, title: "", author: ""})
  const setData = (key: string, value: string) => {
		setBook({ ...book, [key]: value })
	}

  const setEditData = (key: string, value: string) => {
		setEditBook({ ...editBook, [key]: value })
	}

  const checkEmptyValue = async () => {
		type Error = {
			title: boolean,
			author: boolean,
			[key: string]: boolean
		}

		const errCheck: Error = {
			title: false,
			author: false
		}

		errCheck.title = book.title === ""
		errCheck.author = book.author === ""

		return Object.values(errCheck).some(values => values)
	}

  const [open, setOpen] = useState(false);
  
  const handleClose = () => {
    setOpen(false);
  };

  const handleOpen = (book: {id: number, title: string, author: string}) => async () => {
    console.log(book.id);
    
    setEditBook(book)
    setOpen(true);
  };
  
  useEffect(() =>{
    getBooks()
      .then((res)=>{setBooks(res.data)})
      .catch((err)=>{console.error(err)})
  },[])

  const handleCreateBook = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()

		const emptyCheck = await checkEmptyValue()
		if (emptyCheck) {
			return
		}

    axios.request({
      method: "post",
      url: `http://${hostname}:8080/book`,
      data: book
    })
    .then((res)=>{
      if (res.status == 201) {
        getBooks()
          .then((res)=>{setBooks(res.data)})
          .catch((err)=>{console.error(err)})
        setBook({id: 0, title: "", author: ""})
      }
    })
  }

  const handleEditBook = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()
    axios.request({
      method: "patch",
      url: `http://${hostname}:8080/book`,
      data: editBook
    })
    .then((res)=>{
        console.log(res.status)
        getBooks()
          .then((res)=>{setBooks(res.data)})
          .catch((err)=>{console.error(err)})
        setOpen(false);
    })

  }

  const handleDeleteBook = (id: number) => async () => {
    await deleteBook(id)
      .then((res)=>{
        console.log(res.status)
        getBooks()
          .then((res)=>{setBooks(res.data)})
          .catch((err)=>{console.error(err)})
      })
      .catch((err)=>{console.error(err)})
  }


  //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
  const PageLayoutSx: SxProps = {display:"flex", flexDirection:"column", alignItems:"center"}
  const CommandContainerSx: SxProps = {display:"flex", justifyContent:"center", bgcolor: "#ffff", width: 600, padding:5, borderRadius: 3}
  //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
  return (
    <Box sx={PageLayoutSx} >
      <FormCommand handleCreateBook={handleCreateBook} setData={setData} book={book}/>
      {/* Back drop edit book */}
      <Backdrop
        sx={(theme) => ({ color: '#fff', zIndex: theme.zIndex.drawer + 1 })}
        open={open}
      >
        <Container component={'form'} sx={CommandContainerSx} onSubmit={handleEditBook}>
          <Stack direction={"column"} spacing={2}>
            <IconButton aria-label="exit" color='error' onClick={handleClose} sx={{alignSelf: "end",justifySelf: "end"}}>
              <CloseIcon />
            </IconButton>
            <Stack direction={"row"} spacing={2}>
              <TextField id="edit-id" label="ID" variant="outlined" size='small' type='text' sx={{width:60}} value={editBook.id} disabled/>
              <TextField id="edit-title" label="Title" variant="outlined" size='small' type='text' onChange={(e)=>{setEditData("title",e.target.value)}} value={editBook.title}/>
              <TextField id="edit-author" label="Author" variant="outlined" size='small' type='text' onChange={(e)=>{setEditData("author",e.target.value)}} value={editBook.author}/>
              <Button variant="contained" type='submit'>บันทึก</Button>
            </Stack>
          </Stack>
        </Container>
      </Backdrop>
      {books.length == 0?
      ""
      :
      <Stack direction={"column"} spacing={2} alignItems={'center'} sx={{marginBlock:4, height: 600}}>
        {books.map((el: Book, index: number)=>(
          <Stack component={Paper}
            direction={"row"} 
            spacing={4} 
            alignItems={'center'} 
            justifyContent={"space-between"} 
            key={index} 
            height={60}
            width={450}
            paddingInline={3}
            bgcolor={"#292939"}>
            <Typography sx={{color: "#fefefe"}} fontWeight={600}>{`Title: ${el.title}, Author: ${el.author}`}</Typography>
            <Stack direction={"row"} spacing={2} alignItems={'center'} key={index}>
              <IconButton aria-label="edit" color='primary' onClick={handleOpen(el)}>
                <EditIcon />
              </IconButton>
              <IconButton aria-label="delete" color='error' onClick={handleDeleteBook(el.id)}>
                <DeleteIcon />
              </IconButton>
            </Stack>
          </Stack>
        ))}
      </Stack>
      }
      
    </Box>
  )
}

export default App
