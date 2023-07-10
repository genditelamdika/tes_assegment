<h1># tes_assegment</h1>
<h1>#cara menjalankan aplikasi</h1>
<p>#go run main.go</p>
<p>menajalankan di postman</p>
<br>
<h1>Register</h1>
<p>localhost:5000/api/v1/register</p>
<p>{
    "fullname": "",
    "email": "",
    "password": ""
    
}</p>
<h1>Login</h1>

<p>localhost:5000/api/v1/login</p>
<p>{
    "email":"",
    "password":""
}</p>
<h1>CreateProduct</h1>
<p>localhost:5000/api/v1/product</p>
<p>Form data:</p>
<p>name</p>
<p>price</p>
<p>description</p>
<p>image</p>
<p>categoryid</p>

<h1>#untuk Pembelian 
authorizarion User Mengunakan token
http://localhost:5000/api/v1/cart
{
    "productid":,
    "status":""
}
tugas latar belakang yang dijalankan setiap hari pada tengah malam yang dikirim email ke setiap pelanggan dengan pengingat pesanan yang tertunda. 
Email tersebut  menyertakan daftar produk dalam pesanannya dan tautan untuk menyelesaikan proses pembayaran

#melakukan transaction
#kode diskon:
? IC003 : berlaku 10% untuk semua barang 
? IC042 : berlaku 5% untuk barang dengan kategori elektronik 
? IC015 : berlaku 10% untuk semua barang pada hari Sabtu dan Minggu
localhost:5000/api/v1/transaction
{
    "status":"succes",
    "discountcode":""
}

#generate CSV 
http://localhost:5000/api/v1/generatecsv
a. ID Pesanan 
b. Nama pelangganx
c. Tanggal pemesanan 
d. Jumlah harga pesanan 
e. Status pesanan

#membatasi jumlah permintaan per menit dari satu alamat IP menjadi 100
