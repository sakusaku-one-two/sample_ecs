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
        const result = await fetch('/api/',{
            method : "GET"
        });

        const textData = await result.text();
        
        setData(`set get data >>>>> ${textData}`)
      } catch(e:any) {
        setData("error");
      }
    };

    const Post = async () => {
      const res= await fetch('/api/',{
        method : "POST",
        headers : {
          'Content-Type': 'application/json',
        }
      });

      if (!res.ok) {
        console.log("fiald http post request");
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
          Getdata {data} PostData {postData?.name}
          
      </div>
    </>
  )
}

export default App
