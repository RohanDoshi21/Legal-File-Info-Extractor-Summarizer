import React, { useState, useEffect } from "react";
import { storage } from "./firebase.js";
import { ref, uploadBytes } from "firebase/storage";
import { v4 } from "uuid";
import Axios from "axios";
import Navbar from './Navbar.js';

function UserPage() {
  const [selectedPdf, setSelectedPdf] = useState(null);
  const [documents, setDocuments] = useState([]);
  const [jwtToken] = useState(localStorage.getItem("jwtToken"));

  const [editingDocument, setEditingDocument] = useState(null);
  const [showEditWindow, setShowEditWindow] = useState(false);

  const handleEditClick = (document) => {
    setEditingDocument(document);
    setShowEditWindow(true);
  };

  const handleSummaryClick = (document) => {
    const centerX = (window.innerWidth - 500) / 2;
    const centerY = (window.innerHeight - 500) / 2;

    const summaryWindow = window.open(
      "",
      "_blank",
      `width=750,height=500,resizable=yes,scrollbars=yes,left=${centerX},top=${centerY}`
    );

    summaryWindow.document.open();
    summaryWindow.document.write(`
      <html>
        <head>
          <title>Summary of ${document.name}</title>
          <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.16/dist/tailwind.min.css" rel="stylesheet">
        </head>
        <body class="bg-blue-100 p-4">
          <h2 class="text-2xl font-semibold">Summary of ${document.name}</h2>
          <div class="mt-4">
            <textarea readonly class="w-full h-full p-2 border border-blue-500 resize-both rounded-md">${document.content.summary}</textarea>
          </div>
        </body>
      </html>
    `);
    summaryWindow.document.close();
  };

  const renderEditWindow = () => {
    if (!showEditWindow || !editingDocument) return null;

    const handleCancelClick = () => {
      setEditingDocument(null);
      setShowEditWindow(false);
    };

    return (
      <div className="absolute top-0 left-0 w-screen h-screen bg-gray-200 bg-opacity-80 flex items-center justify-center">
        <div className="bg-white p-4 rounded-lg shadow-md w-2/4 h-2/4">
          <h2 className="text-2xl font-semibold mb-4">
            {" "}
            {editingDocument.name} Content
          </h2>
          <div className="max-h-60 overflow-auto">
            <table className="w-full mb-4">
              <tbody>
                {Object.entries(editingDocument.content).map(([key, value], index) => {
                  if (key !== 'summary' && key !== 'pending') {
                    return (
                      <tr key={index}>
                        <td className="w-1/2 text-right pr-2">
                          <input
                            type="text"
                            value={key}
                            className="w-full border border-gray-300 rounded-md p-1 text-lg"
                          />
                        </td>
                        <td className="w-1/2">
                          <input
                            type="text"
                            value={value}
                            className="w-full border border-gray-300 rounded-md p-1 text-lg"
                          />
                        </td>
                      </tr>
                    );
                  }
                  return null; // Don't render a row for 'summary'
                })}
              </tbody>
            </table>
          </div>
          <div className="text-center">
            <button
              onClick={handleCancelClick}
              className="py-2 px-4 bg-indigo-600 hover:bg-indigo-700 text-white rounded-md focus-ring focus-ring-indigo-200"
            >
              Back
            </button>
          </div>
        </div>
      </div>
    );
  };

  const fetchDocuments = async () => {
    try {
      const response = await Axios.get(
        "http://localhost:8080/files",
        {
          headers: {
            Authorization: `Bearer ${jwtToken}`,
          },
        }
      );

      if (response.status === 200) {
        setDocuments(response.data.files);
      } else {
        setDocuments([]); // No documents available
      }
    } catch (error) {
      console.error("Error fetching documents:", error);
    }
  };

  useEffect(() => {
    fetchDocuments();
  }, [jwtToken]); // Fetch documents when the component mounts and whenever jwtToken changes

  const handlePdfDrop = (event) => {
    event.preventDefault();
    const file = event.dataTransfer.files[0];
    if (file && file.type === "application/pdf") {
      setSelectedPdf(file);
    }
  };

  const handlePdfSelect = (event) => {
    const file = event.target.files[0];
    if (file && file.type === "application/pdf") {
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

    // Upload to Firebase Storage
    // const storageRef = ref(storage, `${selectedPdf.name + v4()}`);
    // uploadBytes(storageRef, selectedPdf).then(() => {
    //   alert("Uploaded a file!");
    // });

    const config = {
      headers: {
        "Content-Type": "multipart/form-data",
        Authorization: `Bearer ${jwtToken}`,
      },
    };

    const formData = new FormData();
    formData.append('file', selectedPdf);

    try {
      const response = await Axios.post(
        "http://localhost:8080/files",
        formData,
        config
      );

      if (response.status === 200) {
        alert("Upload successful!");
        // Successful upload, fetch documents again to update the table
        fetchDocuments();
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

  return (
    <>
      <Navbar jwtToken={jwtToken} />
      <div className="bg-blue-200 min-h-screen p-8">
        <h1 className="text-black text-center text-4xl font-bold py-10">
          Welcome User!
        </h1>
        <div
          className="bg-slate-100 text-center mx-96 py-6 border-2 rounded-lg border-black space-y-5"
          onDragOver={(e) => e.preventDefault()}
          onDrop={handlePdfDrop}
        >
          {selectedPdf ? (
            <div>
              <p className="text-2xl bg-slate mb-4">
                PDF Selected: {selectedPdf.name}
              </p>
              <button
                onClick={clearSelectedPdf}
                className="rounded-lg border-2 border-black text-black bg-red-300 px-20 py-2 hover:text-white hover:bg-red-500 font-medium active-bg-red-300 active-text-black"
              >
                Clear PDF
              </button>
            </div>
          ) : (
            <div>
              <p className="text-3xl bg-slate my-4">Drop PDF here</p>
              <p className="my-4">or</p>
              <label className="cursor-pointer rounded-lg border-2 border-black text-black bg-blue-300 px-20 py-2 hover:text-white hover:bg-blue-500 font-medium active-bg-blue-300 active-text-black">
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

        <div className="flex justify-center items-center mt-4">
          <button
            onClick={handleUpload}
            className="p-2 px-8 m-1 bg-indigo-600 hover:bg-indigo-700 text-white rounded-md"
          >
            Upload
          </button>
        </div>

        {documents?.length === 0 ? (
          <p className="text-center text-2xl mt-4">No documents available.</p>
        ) : (
          <>
            <table className="w-full mt-4 bg-white rounded-lg shadow">
              <thead>
                <tr>
                  <th className="border-t-0 border-r-0 border-l-0 border-b border-gray-200 text-center p-3">
                    Name of Document
                  </th>
                  <th className="border-t-0 border-r-0 border-l-0 border-b border-gray-200 text-center p-3">
                    Uploaded
                  </th>
                  <th className="border-t-0 border-r-0 border-l-0 border-b border-gray-200 text-center p-3">
                    Last Modified
                  </th>
                  <th className="border-t-0 border-r-0 border-l-0 border-b border-gray-200 text-center p-3">
                    View Contents
                  </th>
                  <th className="border-t-0 border-r-0 border-l-0 border-b border-gray-200 text-center p-3">
                    Update Status
                  </th>
                  <th className="border-t-0 border-r-0 border-l-0 border-b border-gray-200 text-center p-3">
                  Summary View
                </th>
                </tr>
              </thead>
              <tbody>
                {documents?.map((document, index) => (
                  <tr
                    key={document.id}
                    className={`${index % 2 === 0 ? "bg-blue-100" : "bg-blue-200"
                      } hover:bg-blue-300 transition duration-300`}
                  >
                    <td className="border-t-0 border-r-0 border-l-0 border-b border-gray-200 text-center p-3">
                      {document.name}
                    </td>
                    <td className="border-t-0 border-r-0 border-l-0 border-b border-gray-200 text-center p-3">
                      {new Date(document.created_at).toLocaleString()}
                    </td>
                    <td className="border-t-0 border-r-0 border-l-0 border-b border-gray-200 text-center p-3">
                      {new Date(document.updated_at).toLocaleString()}
                    </td>
                    <td className="border-t-0 border-r-0 border-l-0 border-b border-gray-200 text-center p-3">
                      <button
                        onClick={() => handleEditClick(document)}
                        className="py-2 px-4 bg-indigo-600 hover:bg-indigo-700 text-white rounded-md focus-ring focus-ring-indigo-200"
                      >
                        View
                      </button>
                    </td>
                    <td className="border-t-0 border-r-0 border-l-0 border-b border-gray-200 text-center p-3">
                    <button
                      className={`py-2 px-4 ${document.content?.pending === true
                        ? "bg-red-600 hover:bg-red-700"
                        : "bg-green-600 hover:bg-green-700"
                        } text-white rounded-md focus-ring focus-ring-indigo-200`}
                    >
                      {document.content?.pending ? "Pending" : "Updated"}
                    </button>
                  </td>
                  <td className="border-t-0 border-r-0 border-l-0 border-b border-gray-200 text-center p-3">
                                    <button
                                        onClick={() => handleSummaryClick(document)}
                                        className="py-2 px-4 bg-indigo-600 hover:bg-indigo-700 text-white rounded-md focus:ring focus:ring-indigo-200"
                                    >
                                        Summary
                                    </button>
                                </td>
                  </tr>
                ))}
              </tbody>

            </table>
            {renderEditWindow()}
          </>
        )}
      </div>
    </>
  );
}

export default UserPage;
