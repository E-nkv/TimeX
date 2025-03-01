export async function fetchCategories(){
    const res = await fetch("http://localhost:8080/categories")
    if (!res.ok) {
        alert("Error fetching the categories: " +  res.status)
        return
    }
    const d = await res.json()
    return d.categories
    
}