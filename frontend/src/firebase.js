// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import {getStorage} from 'firebase/storage';
 

// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
const firebaseConfig = {
  apiKey: "AIzaSyBVscdvf7s1NN5YdM9-BH_Vj5aTD07z_xw",
  authDomain: "wordextract-e039a.firebaseapp.com",
  projectId: "wordextract-e039a",
  storageBucket: "wordextract-e039a.appspot.com",
  messagingSenderId: "658685854166",
  appId: "1:658685854166:web:7a7f3010b3d469df9fa477"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
export const storage = getStorage(app);