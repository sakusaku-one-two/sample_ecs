import { useState,useEffect } from 'react';
import { Button } from "@/components/ui/button"

import './App.css'

interface ResultType {
  name :string
}

function App() {
  const [add,setAdd] = useState<string>("");
  const [data, setData] = useState<string[]>([]);
  const [postData,setPostData] = useState<ResultType|undefined>(undefined);




  useEffect(()=>{
    const get = async () => {
      try {
        const result = await fetch('/api/all',{
            method : "GET",
        });

        const textData = await result.json();
        setData(textData.result)
      } catch(e:any) {
        setData(["error"]);
      }
    };
    get()
    
  },[]);

  const Post = async () => {
    try {
      const res = await fetch('/api/add', {
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
      <div className='border-gray-300 rounded bg-gray-100'>
          Getdata {data} requestURL is {postData?.name}
          <input type="text" onChange={(e)=> setAdd(e.target.value) } value={add}/>
          <Button variant="outline" onClick={() => Post()}>POST</Button>
          
      </div>
    </>
  )
}

export default App
