import { useState,useEffect } from 'react';
import { Button } from "@/components/ui/button"

import './App.css'

interface ResultType {
  name :string
}

const redisChecker = async (callback:Function) => {
  try {
    const reponse = await fetch("/api/redis");
    const data = await reponse.json();
    callback(data);
  } catch (e:any) {
    callback(e.message as string);
  }
};

const get = async (setterCallBack:Function) => {
  try {
    const result = await fetch('/api/all',{
        method : "GET",
    });

    const textData = await result.json();
    setterCallBack(textData.result)
  } catch(e:any) {
    setterCallBack(["error"]);
  }
};

function App() {
  const [reidsState,setRedisState] = useState<string>("");
  const [add,setAdd] = useState<string>("");
  const [data, setData] = useState<string[]>([]);
  const [postData,setPostData] = useState<ResultType|undefined>(undefined);




  useEffect(()=>{
    
    get(setData)
    
  },[]);

  const Post = async () => {
    try {
      const res = await fetch('/api/add', {
        method: "POST",
        headers: {
          'Content-Type': 'application/json',

        },
        body:JSON.stringify({value:add})
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
         
          <input type="text" onChange={(e)=> setAdd(e.target.value) } value={add}/>
          <Button variant="outline" onClick={() =>{
            Post();
            setAdd("");
          } }>POST</Button>
          {reidsState}
          <Button onClick={() => redisChecker(setRedisState)} />
            <div className='flex flex-col'>
                {
                  data.map((value,index) => {
                    return (
                      <div key={index} className='rounded bg-gray-300'>
                          {value}
                      </div>
                    )
                  })
                }
                {
                  postData?.name
                }
            </div>
         
          
      </div>
    </>
  )
}

export default App
