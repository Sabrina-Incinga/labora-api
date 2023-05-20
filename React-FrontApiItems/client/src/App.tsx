import { useEffect, useState } from 'react'
import './App.css'

async function getItems(page: number, itemsPerPage:number) {
  const res = await fetch(`http://localhost:8000/items?page=${page}&itemsPerPage=${itemsPerPage}`)
  const data = await res.json()

  return await data
}



function App() {
  const [page, setPage] = useState(0)
  const [totalItems, setTotalItems] = useState(0)
  const [items, setItems] = useState([{
    id: 0,
    name: '',
    order_date: Date.now(),
    product: '',
    quantity: 0,
    price: 0,
    details: '',
    total_price: 0
  }]);

  

  useEffect(()=>{

    getItems(page, 100).then(data => {
      setItems(data.Items)
      setTotalItems(data.ItemCount)
    })
    
  }, [page])

  console.log(items);

  return (
    <>
      {items?.map(item => (
        <div key={item.id}>
          <h2>{item.name}</h2>
          <div>
            <div>{item.product}</div>
            <div>{item.quantity}</div>
          </div>
        </div>
      ))}
      <button type='button' onClick={() => setPage(prevState => prevState-1)} disabled={page == 0}>Prev</button>
      <button type='button' onClick={() => setPage(prevState => prevState+1)} disabled={totalItems < (page+1)*100}>Next</button>
    </>
  )
}

export default App
