->project structure
-main → start
-cmd → server logic
-handlers → business logic
-middleware → request processing
-database → data

->Full System Flow

    Client (React/Postman)
            ↓
    Preflight (OPTIONS check)
             ↓
    Cors (allow request)
             ↓
    Logger (track request)
             ↓
    Router (ServeMux)
             ↓
    Handler (Get/Add Product)
             ↓
    SendData (response)
             ↓
          Client
