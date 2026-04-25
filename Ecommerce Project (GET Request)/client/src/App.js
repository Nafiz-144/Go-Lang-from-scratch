import React, { useEffect, useState } from "react";

function App() {
  const [products, setProducts] = useState([]);
  const [form, setForm] = useState({
    title: "",
    discription: "",
    price: "",
    imgUrl: "",
  });

  // Fetch products
  const getProducts = async () => {
    try {
      const res = await fetch("http://localhost:8080/getproduct");
      const data = await res.json();
      setProducts(data);
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    getProducts();
  }, []);

  // Handle input change
  const handleChange = (e) => {
    setForm({
      ...form,
      [e.target.name]: e.target.value,
    });
  };

  // Submit form
  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const res = await fetch("http://localhost:8080/addproduct", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          title: form.title,
          discription: form.discription,
          price: parseFloat(form.price),
          imgUrl: form.imgUrl,
        }),
      });

      const data = await res.json();
      console.log(data);

      // Refresh list
      getProducts();

      // Clear form
      setForm({
        title: "",
        discription: "",
        price: "",
        imgUrl: "",
      });
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div style={{ padding: "20px" }}>
      <h1>Product App</h1>

      {/* FORM */}
      <form onSubmit={handleSubmit}>
        <input
          name="title"
          placeholder="Title"
          value={form.title}
          onChange={handleChange}
          required
        />
        <br />

        <input
          name="discription"
          placeholder="Description"
          value={form.discription}
          onChange={handleChange}
          required
        />
        <br />

        <input
          name="price"
          placeholder="Price"
          type="number"
          value={form.price}
          onChange={handleChange}
          required
        />
        <br />

        <input
          name="imgUrl"
          placeholder="Image URL"
          value={form.imgUrl}
          onChange={handleChange}
          required
        />
        <br />

        <button type="submit">Add Product</button>
      </form>

      <hr />

      {/* PRODUCT LIST */}
      <h2>Products</h2>
      {products.map((p) => (
        <div
          key={p.id}
          style={{ border: "1px solid black", margin: "10px", padding: "10px" }}
        >
          <h3>{p.title}</h3>
          <p>{p.discription}</p>
          <p>Price: {p.price}</p>
          <img src={p.imgUrl} alt="" width="150" />
        </div>
      ))}
    </div>
  );
}

export default App;
