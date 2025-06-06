import { useState,useEffect } from 'react'

import './App.css'

interface ResultType {
  name :string
}

function App() {
  const [data, setData] = useState<string>("get dta");
  const [postData,setPostData] = useState<ResultType|undefined>(undefined);




  useEffect(()=>{
    const get = async () => {
      try {
        const result = await fetch(`${window.location.host}/api/`,{
            method : "GET"
        });

        const textData = await result.text();
        setData(textData)
      } catch(e:any) {
        setData("error");
      }
    };

    const Post = async () => {
      const res= await fetch(`${window.location.host}/api/`,{
        method : "POST",
        headers : {
          'Content-Type': 'application/json',
        }
      });

      if (!res.ok) {
        throw new Error("faild http request");

      }

      const data:ResultType = await res.json();

      setPostData(data);

    };
  

    get()
    Post();
  },[]);


  return (
    <>
      <div className=''>
          Getdata {data} PostData {postData}
          
      </div>
    </>
  )
}

export default App
