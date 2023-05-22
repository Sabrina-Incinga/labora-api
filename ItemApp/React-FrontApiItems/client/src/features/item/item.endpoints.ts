import { itemDTO } from "../types/item.types"

const endpoint = 'http://localhost:8000/items'

export async function getItems(page: number, itemsPerPage:number) {
    const res = await fetch(`${endpoint}?page=${page}&itemsPerPage=${itemsPerPage}`)
    const data = await res.json()
  
    return await data
}

export async function updateItemRequest(id: number, payload: itemDTO) {
  const config = {
    method: "PUT",
    body: JSON.stringify(payload),
    withCredentials: true,
    crossdomain: true,
    headers: { "Content-Type": "application/json" },
 }
  const res = await fetch(`${endpoint}/update/${id}`, config)
  const data = await res.text()

  return await data
}

export async function deleteItemRequest(id: number) {
  const config = {
    method: "DELETE",
    withCredentials: true,
    crossdomain: true,
    headers: { "Content-Type": "application/json" },
  }

  const res = await fetch(`${endpoint}/delete/${id}`,config)
  const data = await res.text()

  return await data
}