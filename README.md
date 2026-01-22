# Kasir API

RESTful API untuk manajemen produk dan kategori kasir.

**Base URL:** `https://kasir-api-production-6770.up.railway.app`

---

## API Reference

### Health Check

#### Get health status

```http
GET /health
```

Cek status API apakah sedang berjalan.

---

### Products

#### Get all products

```http
GET /api/products
```

Mengambil semua data produk.

---

#### Get product by ID

```http
GET /api/products/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int`    | **Required**. ID produk           |

Mengambil detail produk berdasarkan ID.

---

#### Create product

```http
POST /api/products
```

| Body Field | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`    | `string` | **Required**. Nama produk         |
| `price`   | `int`    | **Required**. Harga (harus > 0)   |
| `stock`   | `int`    | **Required**. Stok (harus > 0)   |

Membuat produk baru.

---

#### Update product

```http
PUT /api/products/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int`    | **Required**. ID produk           |

| Body Field | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`    | `string` | **Required**. Nama produk         |
| `price`   | `int`    | **Required**. Harga (harus > 0)   |
| `stock`   | `int`    | **Required**. Stok (harus > 0)   |

Mengupdate data produk berdasarkan ID.

---

#### Delete product

```http
DELETE /api/products/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int`    | **Required**. ID produk           |

Menghapus produk berdasarkan ID.

---

### Categories

#### Get all categories

```http
GET /api/categories
```

Mengambil semua data kategori.

---

#### Get category by ID

```http
GET /api/categories/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int`    | **Required**. ID kategori         |

Mengambil detail kategori berdasarkan ID.

---

#### Create category

```http
POST /api/categories
```

| Body Field     | Type     | Description                       |
| :------------- | :------- | :-------------------------------- |
| `name`         | `string` | **Required**. Nama kategori       |
| `description`  | `string` | Optional. Deskripsi kategori      |

Membuat kategori baru.

---

#### Update category

```http
PUT /api/categories/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int`    | **Required**. ID kategori         |

| Body Field     | Type     | Description                       |
| :------------- | :------- | :-------------------------------- |
| `name`         | `string` | **Required**. Nama kategori       |
| `description`  | `string` | Optional. Deskripsi kategori      |

Mengupdate data kategori berdasarkan ID.

---

#### Delete category

```http
DELETE /api/categories/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int`    | **Required**. ID kategori         |

Menghapus kategori berdasarkan ID.

---

## Response Schema

### Success Response

```json
{
  "success": true,
  "message": "Success message",
  "data": {} // atau [] untuk array, atau null
}
```

### Error Response

```json
{
  "success": false,
  "message": "Error message",
  "data": null
}
```

### HTTP Status Codes

- `200 OK` - Request berhasil
- `201 Created` - Resource berhasil dibuat
- `400 Bad Request` - Request tidak valid (validation error)
- `404 Not Found` - Resource tidak ditemukan

---

## Example Usage

### Health Check

```bash
curl https://kasir-api-production-6770.up.railway.app/health
```

### Products

#### Get all products

```bash
curl https://kasir-api-production-6770.up.railway.app/api/products
```

#### Get product by ID

```bash
curl https://kasir-api-production-6770.up.railway.app/api/products/1
```

#### Create product

```bash
curl -X POST https://kasir-api-production-6770.up.railway.app/api/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Kopi",
    "price": 3000,
    "stock": 10
  }'
```

#### Update product

```bash
curl -X PUT https://kasir-api-production-6770.up.railway.app/api/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Kopi Hitam",
    "price": 3500,
    "stock": 15
  }'
```

#### Delete product

```bash
curl -X DELETE https://kasir-api-production-6770.up.railway.app/api/products/1
```

### Categories

#### Get all categories

```bash
curl https://kasir-api-production-6770.up.railway.app/api/categories
```

#### Get category by ID

```bash
curl https://kasir-api-production-6770.up.railway.app/api/categories/1
```

#### Create category

```bash
curl -X POST https://kasir-api-production-6770.up.railway.app/api/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Snack",
    "description": "Kategori untuk snack"
  }'
```

#### Update category

```bash
curl -X PUT https://kasir-api-production-6770.up.railway.app/api/categories/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Makanan Ringan",
    "description": "Kategori untuk makanan ringan"
  }'
```

#### Delete category

```bash
curl -X DELETE https://kasir-api-production-6770.up.railway.app/api/categories/1
```

---

## Notes

- API menggunakan in-memory storage (data akan hilang saat server restart)
- Semua endpoint mengembalikan response dalam format JSON
- Validation dilakukan untuk semua field required
- ID di-generate secara otomatis (auto-increment)
