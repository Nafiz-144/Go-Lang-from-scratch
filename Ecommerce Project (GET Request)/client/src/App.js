import React, { useEffect, useState } from "react";

function App() {
  const [products, setProducts] = useState([]);

  const [form, setForm] = useState({
    title: "",
    description: "",
    price: "",
    imgUrl: "",
  });

  const [searchId, setSearchId] = useState("");
  const [singleProduct, setSingleProduct] = useState(null);

  // ================= GET ALL PRODUCTS =================
  const getProducts = async () => {
    try {
      const res = await fetch("http://localhost:8080/products");
      const data = await res.json();
      setProducts(data);
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    getProducts();
  }, []);

  // ================= HANDLE INPUT =================
  const handleChange = (e) => {
    setForm({
      ...form,
      [e.target.name]: e.target.value,
    });
  };

  // ================= ADD PRODUCT =================
  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const res = await fetch("http://localhost:8080/products", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          title: form.title,
          description: form.description,
          price: parseFloat(form.price),
          imgUrl: form.imgUrl,
        }),
      });

      const data = await res.json();
      console.log(data);

      // refresh list
      getProducts();

      // reset form
      setForm({
        title: "",
        description: "",
        price: "",
        imgUrl: "",
      });
    } catch (err) {
      console.error(err);
    }
  };

  // ================= GET PRODUCT BY ID =================
  const getProductById = async () => {
    try {
      const res = await fetch(`http://localhost:8080/product/${productId}`);

      if (!res.ok) {
        throw new Error("Product not found");
      }

      const data = await res.json();
      setSingleProduct(data);
    } catch (err) {
      console.error(err);
      setSingleProduct(null);
      alert("Product not found");
    }
  };

  return (
    <div style={{ padding: "20px" }}>
      <h1>Product App</h1>

      {/* ================= FORM ================= */}
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
          name="description"
          placeholder="Description"
          value={form.description}
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

      {/* ================= SEARCH BY ID ================= */}
      <hr />

      <h2>Search Product by ID</h2>

      <input
        type="number"
        placeholder="Enter Product ID"
        value={searchId}
        onChange={(e) => setSearchId(e.target.value)}
      />

      <button onClick={getProductById}>Search</button>

      {/* SINGLE PRODUCT */}
      {singleProduct && (
        <div
          style={{
            border: "2px solid green",
            margin: "10px",
            padding: "10px",
          }}
        >
          <h3>{singleProduct.title}</h3>
          <p>{singleProduct.description}</p>
          <p>Price: {singleProduct.price}</p>
          <img src={singleProduct.imgUrl} alt="" width="150" />
        </div>
      )}

      {/* ================= ALL PRODUCTS ================= */}
      <hr />

      <h2>All Products</h2>

      {products.map((p) => (
        <div
          key={p.id}
          style={{
            border: "1px solid black",
            margin: "10px",
            padding: "10px",
          }}
        >
          <h3>{p.title}</h3>
          <p>{p.description}</p>
          <p>Price: {p.price}</p>
          <img src={p.imgUrl} alt="" width="150" />
        </div>
      ))}
    </div>
  );
}

export default App;
