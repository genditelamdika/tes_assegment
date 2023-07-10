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
<p>authorizarion User Mengunakan token</p>
<p>http://localhost:5000/api/v1/cart</p>
<p>{
    "productid":,
    "status":""
}</p>
<p>tugas latar belakang yang dijalankan setiap hari pada tengah malam yang dikirim email ke setiap pelanggan dengan pengingat pesanan yang tertunda. 
Email tersebut  menyertakan daftar produk dalam pesanannya dan tautan untuk menyelesaikan proses pembayaran
</p>
<h1>#melakukan transaction</h1>
<h2>#kode diskon:</h2>
<p>? IC003 : berlaku 10% untuk semua barang </p>
<p>? IC042 : berlaku 5% untuk barang dengan kategori elektronik</p> 
<p>? IC015 : berlaku 10% untuk semua barang pada hari Sabtu dan Minggu</p>
<p>localhost:5000/api/v1/transaction</p>
<p>{
    "status":"succes",
    "discountcode":""
}</p>

<h1>#generate CSV</h1> 
<p>http://localhost:5000/api/v1/generatecsv</p>
<p>a. ID Pesanan</p> 
<p>b. Nama pelangganx</p>
<p>c. Tanggal pemesanan</p> 
<p>d. Jumlah harga pesanan</p> 
<p>e. Status pesanan</p>

#membatasi jumlah permintaan per menit dari satu alamat IP menjadi 100
