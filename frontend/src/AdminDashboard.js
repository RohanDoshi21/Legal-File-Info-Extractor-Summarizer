import { useNavigate } from 'react-router-dom';
import React, { useState, useEffect } from 'react';
import Axios from "axios";
import Navbar from './Navbar.js';

const AdminPage = () => {
  const navigate = useNavigate();
  const [jwtToken] = useState(localStorage.getItem("jwtToken"));
  const [users, setUsers] = useState([]);
  const [selectedPdf, setSelectedPdf] = useState(null);
  const [adminId, setAdminId] = useState(null);
  const [adminemail, setAdminEmail] = useState(null);

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const response = await Axios.get("http://localhost:8080/admin/users", {
          headers: {
            Authorization: `Bearer ${jwtToken}`,
          },
        });

        if (response.data.ok === 1 && response.data.users) {
          const nonAdminUsers = response.data.users.filter(user => user.isadmin === false);
          setUsers(nonAdminUsers);
          const adminUser = response.data.users.find(user => user.isadmin === true);
          if (adminUser) {
            setAdminId(adminUser.id);
            setAdminEmail(adminUser.email);
          }
        }
      } catch (error) {
        console.error("Error fetching users:", error);
      }
    };

    fetchUsers();
  }, [jwtToken]); // Fetch users when the component mounts and whenever jwtToken changes


  const handlePdfDrop = (event) => {
    event.preventDefault();
    const file = event.dataTransfer.files[0];
    if (file && file.type === 'application/pdf') {
      setSelectedPdf(file);
    }
  };

  const handlePdfSelect = (event) => {
    const file = event.target.files[0];
    if (file && file.type === 'application/pdf') {
      setSelectedPdf(file);
    }
  };

  const clearSelectedPdf = () => {
    setSelectedPdf(null);
  };

  const handleUpload = async () => {
    if (!selectedPdf) {
      alert("No PDF chosen. Please select a PDF file before uploading.");
      return;
    }

    const config = {
      headers: {
        "Content-Type": "multipart/form-data",
        Authorization: `Bearer ${jwtToken}`,
      },
    };

    const formData = new FormData();
    formData.append('file', selectedPdf);

    try {
      // const newDocument = {
      //   file_name: selectedPdf.name,
      //   file_content: {
      //     summary: "Summary of the uploaded document",
      //     sample_key1: "sample_value_admin1",
      //     sample_key2: "sample_value_admin2",
      //   },
      //   file_link: "sample-linkk",
      // };
      const response = await Axios.post(
        "http://localhost:8080/files",
        formData,
        config
      );

      if (response.status === 200) {
        alert("Upload successful!");
        // Successful upload, fetch documents again to update the table
      } else {
        // Handle upload error
        console.error("Upload error:", response.data.message || "An error occurred during upload.");
      }
    } catch (error) {
      console.error("Network error during upload:", error);
    } finally {
      clearSelectedPdf();
    }
  };

  const handleUserClick = (user) => {
    navigate('/AdminViewInfo', { state: { userEmail: user.email, userId: user.id } });
  };

  return (
    <>
     <Navbar jwtToken={jwtToken} />
      <div className="p-4 bg-gradient-to-b from-blue-200 via-blue-300 to-blue-200">
        <h2 className="text-3xl text-center text-gray-800 font-semibold mb-4">Upload Document</h2>
        <div
          className="bg-slate-100 text-center mx-96 py-6 border-2 rounded-lg border-black space-y-5"
          onDragOver={(e) => e.preventDefault()}
          onDrop={handlePdfDrop}
        >
          {selectedPdf ? (
            <div>
              <p className="text-2xl bg-slate mb-4">PDF Selected: {selectedPdf.name}</p>
              <button
                onClick={clearSelectedPdf}
                className="rounded-lg border-2 border-black text-black bg-red-300 px-20 py-2 hover:text-white hover:bg-red-500 font-medium active-bg-red-300 active:text-black"
              >
                Clear PDF
              </button>
            </div>
          ) : (
            <div>
              <p className="text-3xl bg-slate my-4">Drop PDF here</p>
              <p className="my-4">or</p>
              <label className="cursor-pointer rounded-lg border-2 border-black text-black bg-blue-300 px-20 py-2 hover:text-white hover:bg-blue-500 font-medium active-bg-blue-300 active:text-black">
                Select PDF from your device
                <input
                  type="file"
                  accept=".pdf"
                  className="hidden"
                  onChange={handlePdfSelect}
                />
              </label>
            </div>
          )}
        </div>

        <div className="flex flex-col items-center mt-4">
          <button
            onClick={handleUpload}
            className="p-2 px-20 m-1 bg-indigo-600 hover:bg-indigo-700 focus:ring focus:ring-indigo-200 text-white rounded-md"
          >
            Upload
          </button>
          <button
            className="p-2 px-2 m-1 bg-indigo-600 hover:bg-indigo-700 focus:ring focus:ring-indigo-200 text-white rounded-md"
            onClick={() => navigate('/AdminViewInfo', { state: { userEmail: adminemail, userId: adminId } })}
          >
            View Uploaded Documents
          </button>
        </div>

        <div className="mt-4 px-4">

        <h1 className="text-3xl my-4 mt-10 font-semibold text-center mb-6 text-gray-800">Users List</h1>


        <table className="w-full px-72 mt-4 bg-white rounded-lg shadow">
          <thead>
            <tr>
              <th className="border-t-0 border-r-0 border-l-0 border-b border-gray-200 text-center p-3">
                User Name
              </th>
            </tr>
          </thead>
          <tbody>
            {users.map((user, index) => (
              <tr
                key={user.email}
                onClick={() => handleUserClick(user)}
                className={`cursor-pointer ${index % 2 === 0 ? 'bg-blue-100' : 'bg-blue-200'
                  } hover:bg-blue-300 transition duration-100`}
              >
                <td className="border-t-0 border-r-0 border-l-0 border-b border-gray-200 text-center p-3">
                  {user.email}
                </td>
              </tr>
            ))}
          </tbody>

        </table>
        </div>
      </div>
    </>
  );
};

export default AdminPage;
