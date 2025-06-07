import { useState,useEffect } from 'react';
import { Button } from "@/components/ui/button"

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
    get()
    
  },[]);

  const Post = async () => {
    try {
      const res = await fetch('/api/', {
        method: "POST",
        headers: {
          'Content-Type': 'application/json',
        }
      });

      if (!res.ok) {
        console.log(`HTTP Error: ${res.status} ${res.statusText}`);
        return; // エラー時は処理を中断
      }

      const data: ResultType = await res.json();
      console.log(data);
      setPostData(data);

    } catch (error) {
      console.error('POST request failed:', error);
    }
  };



  return (
    <>
      <div className=''>
          Getdata {data} PostData {postData?.name}
          <Button variant="outline" onClick={() => Post()}>POST</Button>
          
      </div>
    </>
  )
}

export default App
