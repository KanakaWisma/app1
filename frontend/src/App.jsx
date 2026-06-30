import { useEffect, useState } from "react";
import api from "./api";
import "./App.css";

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
  api.post("/mahasiswa", form)
    .then(() => {
      loadMahasiswa();

      alert("Data berhasil ditambah");

      setForm({
        Nama: "",
        Umur: 0,
        NIM: "",
        TglLahir: "",
        Alamat: "",
        IDJurusan: 0,
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

const [selectedId, setSelectedId] = useState(null);

const handleUpdate = () => {
  if (!selectedId) {
    alert("Pilih data terlebih dahulu");
    return;
  }

  api.put(`/mahasiswa/${selectedId}`, form)
    .then(() => {
      loadMahasiswa();
      alert("Data berhasil diupdate");
    })
    .catch((err) => {
      console.log(err);
    });
};

const handleDelete = () => {
  if (!selectedId) {
    alert("Pilih data terlebih dahulu");
    return;
  }

  api.delete(`/mahasiswa/${selectedId}`)
    .then(() => {
      loadMahasiswa();

      setSelectedId(null);

      setForm({
        Nama: "",
        Umur: 0,
        NIM: "",
        TglLahir: "",
        Alamat: "",
        IDJurusan: 0,
      });

      alert("Data berhasil dihapus");
    })
    .catch((err) => {
      console.log(err);
    });
};

  return (
  <div className="container">
    <h1>Data Mahasiswa</h1>

    <div className="content">

      <div className="form-section">

        <label>Nama</label>
        <input
          value={form.Nama}
          onChange={(e) =>
            setForm({ ...form, Nama: e.target.value })
          }
        />

        <label>Umur</label>
        <input
          value={form.Umur}
          onChange={(e) =>
            setForm({
              ...form,
              Umur: Number(e.target.value),
            })
          }
        />

        <label>NIM</label>
        <input
          value={form.NIM}
          onChange={(e) =>
            setForm({ ...form, NIM: e.target.value })
          }
        />

        <label>Tanggal Lahir</label>
        <input
          type="date"
          value={form.TglLahir}
          onChange={(e) =>
            setForm({
              ...form,
              TglLahir: e.target.value,
            })
          }
        />

        <label>Alamat</label>
        <textarea
          value={form.Alamat}
          onChange={(e) =>
            setForm({
              ...form,
              Alamat: e.target.value,
            })
          }
        />

        <label>Jurusan</label>
        <select
          value={form.IDJurusan}
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

      </div>

      <div className="list-section">

        <input
          className="search"
          placeholder="Cari mahasiswa..."
        />

        <table>
          <thead>
            <tr>
              <th>Nama</th>
              <th>Jurusan</th>
            </tr>
          </thead>

          <tbody>
            {mahasiswa.map((mhs) => (
              <tr
      key={mhs.ID}
      onClick={() => {
        setSelectedId(mhs.ID);

        setForm({
          Nama: mhs.Nama,
          Umur: mhs.Umur,
          NIM: mhs.NIM,
          TglLahir: mhs.TglLahir.split("T")[0],
          Alamat: mhs.Alamat,
          IDJurusan: mhs.IDJurusan,
        });
      }}
      style={{ cursor: "pointer" }}
    >
      <td>{mhs.Nama}</td>
      <td>{mhs.Jurusan?.NamaJurusan}</td>
    </tr>
  ))}
          </tbody>
        </table>

      </div>

    </div>

    <div className="action-buttons">

  <button
    onClick={() => {
      setSelectedId(null);

      setForm({
        Nama: "",
        Umur: 0,
        NIM: "",
        TglLahir: "",
        Alamat: "",
        IDJurusan: 0,
      });
    }}
  >
    Reset
  </button>

  <button onClick={handleUpdate}>
    Update
  </button>

  <button onClick={handleSubmit}>
    Save
  </button>

  <button onClick={handleDelete}>
    Delete
  </button>

</div>

  </div>
);
}

export default App;