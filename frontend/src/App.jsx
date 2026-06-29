import { useEffect, useState } from "react";
import api from "./api";

function App() {
  const [mahasiswa, setMahasiswa] = useState([]);
  const [jurusan, setJurusan] = useState([]);

  const [form, setForm] = useState({
  Nama: "",
  Umur: "",
  NIM: "",
  TglLahir: "",
  Alamat: "",
  IDJurusan: "",
});

  useEffect(() => {
  loadMahasiswa();

  api.get("/jurusan")
    .then((res) => {
      setJurusan(res.data.data);
    });
}, []);

const handleSubmit = () => {
  
  alert(JSON.stringify(form));
  
  api.post("/mahasiswa", form)
    .then(() => {

      loadMahasiswa();

      alert("Data berhasil ditambah");

      setForm({
        Nama: "",
        Umur: "",
        NIM: "",
        TglLahir: "",
        Alamat: "",
        IDJurusan: "",
      });

    })
    .catch((err) => {
      console.log(err);
    });
};

const loadMahasiswa = () => {
  api.get("/mahasiswa")
    .then((res) => {
      setMahasiswa(res.data.data);
    });
};

  return (
    <div>
      <h1>Data Mahasiswa</h1>

       <input
      placeholder="Nama"
      onChange={(e) =>
        setForm({ ...form, Nama: e.target.value })
      }
    />

    <br /><br />

    <input
      placeholder="Umur"
      onChange={(e) =>
        setForm({ ...form, Umur: Number(e.target.value) })
      }
    />

    <br /><br />

    <input
      placeholder="NIM"
      onChange={(e) =>
        setForm({ ...form, NIM: e.target.value })
      }
    />

    <br /><br />

    <input
  type="date"
  onChange={(e) =>
    setForm({ ...form, TglLahir: e.target.value })
  }
/>

<br /><br />

<textarea
  placeholder="Alamat"
  onChange={(e) =>
    setForm({ ...form, Alamat: e.target.value })
  }
/>

<br /><br />

    <br /><br />

   <select
  onChange={(e) =>
    setForm({
      ...form,
      IDJurusan: Number(e.target.value),
    })
  }
>
  <option value="">Pilih Jurusan</option>

  {jurusan.map((jrs) => (
    <option
      key={jrs.IDJurusan}
      value={jrs.IDJurusan}
    >
      {jrs.NamaJurusan}
    </option>
  ))}
</select>

<br /><br />

<button onClick={handleSubmit}>
  Save
</button>

    <br /><br />

      {mahasiswa.map((mhs) => (
        <div key={mhs.ID}>
          {mhs.Nama} - {mhs.Jurusan.NamaJurusan}
        </div>
      ))}
    </div>
  );
}

export default App;