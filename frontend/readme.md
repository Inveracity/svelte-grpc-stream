# Frontend

## Dev

```sh
# Install dependencies
cd frontend
npm install --include=dev

# generate gRPC code
cd ..
make proto

# Run
npm run dev -- --open
```
