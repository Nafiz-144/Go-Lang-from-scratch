import React, { useEffect, useState } from "react";

const API = "http://localhost:8080";

function App() {
  const [products, setProducts] = useState([]);
  const [form, setForm] = useState({
    id: "",
    Title: "",
    Description: "",
    Price: "",
    ImgUrl: "",
  });

  // GET
  const fetchProducts = async () => {
    const res = await fetch(`${API}/getproduct`);
    const data = await res.json();
    setProducts(data);
  };

  useEffect(() => {
    fetchProducts();
  }, []);

  // HANDLE INPUT
  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  // POST
  const addProduct = async () => {
    await fetch(`${API}/product`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(form),
    });
    fetchProducts();
  };

  // PUT (full update)
  const updateProduct = async () => {
    await fetch(`${API}/product/${form.id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(form),
    });
    fetchProducts();
  };

  // PATCH (partial update)
  const patchProduct = async () => {
    await fetch(`${API}/product/${form.id}`, {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        Title: form.Title,
      }),
    });
    fetchProducts();
  };

  // DELETE
  const deleteProduct = async (id) => {
    await fetch(`${API}/product/${id}`, {
      method: "DELETE",
    });
    fetchProducts();
  };

  return (
    <div style={{ padding: "20px" }}>
      <h2>Product CRUD</h2>

      {/* FORM */}
      <input name="id" placeholder="ID" onChange={handleChange} />
      <input name="Title" placeholder="Title" onChange={handleChange} />
      <input
        name="Discription"
        placeholder="Description"
        onChange={handleChange}
      />
      <input name="Price" placeholder="Price" onChange={handleChange} />
      <input name="ImgUrl" placeholder="Image URL" onChange={handleChange} />

      <br />
      <br />

      <button onClick={addProduct}>POST</button>
      <button onClick={updateProduct}>PUT</button>
      <button onClick={patchProduct}>PATCH</button>

      <hr />

      {/* LIST */}
      {products.map((p) => (
        <div
          key={p.id}
          style={{ border: "1px solid gray", margin: "10px", padding: "10px" }}
        >
          <h3>{p.Title}</h3>
          <p>{p.Discription}</p>
          <p>Price: {p.Price}</p>
          <img src={p.ImgUrl} width="100" alt="" />
          <br />
          <button onClick={() => deleteProduct(p.id)}>DELETE</button>
        </div>
      ))}
    </div>
  );
}

export default App;
