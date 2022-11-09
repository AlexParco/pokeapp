import { RequireAuth } from '@/components'
import { AuthProvider } from '@/context/auth.context'
import { PokeProvider } from '@/context/pokemon.context'
import { Home, Login, PokemonProfile, Register } from '@/pages'
import { Button, ChakraProvider } from '@chakra-ui/react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css'

function App() {

  return (
    <ChakraProvider>
      <BrowserRouter>
        <AuthProvider>
          <PokeProvider>
            <Routes>
              <Route path="/"
                element={
                  <RequireAuth>
                    <Home />
                  </RequireAuth>
                } />
              <Route path="/:id" element={
                <RequireAuth>
                  <PokemonProfile />
                </RequireAuth>
              } />
              <Route path="/login" element={<Login />} />
              <Route path="/register" element={<Register />} />
            </Routes>
          </PokeProvider>
        </AuthProvider>
      </BrowserRouter>
    </ChakraProvider >
  )
}

export default App
