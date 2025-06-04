import { useState,useEffect } from 'react'

import './App.css'

function App() {
  const [data, setData] = useState<string>("get dta");



  useEffect(()=>{
    const get = async () => {
      try {
        const result = await fetch(`${window.location.host}/api/`,{
            method : "GET"
        });

        const textData = await result.text();
        setData(textData)
      } catch(e:any) {
        setData(e.message);
      }
    };
  

    get()
  },[]);


  return (
    <>
      <div className=''>
          POSTdata {data}
          
      </div>
    </>
  )
}

export default App
