import { Button } from "./ui/button"
import { Timer } from "./timer"
import { useRef, useState, useEffect } from "react";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "./ui/select";
import { Category } from "@/types/types";
import { fetchCategories } from "@/common";


export function TimerComponent(){
    
    const [secs, setSecs] = useState(0);
    const intervalRef = useRef<NodeJS.Timeout>(undefined);
    const [isRunning, setIsRunning] = useState(false)
    const startTimeRef = useRef<number>(0)
    const categoryRef = useRef("1")
    const focusRef = useRef("4")
    const [categories, setCategories] = useState<Category[]>([]); // State for categories

    useEffect(() => {
        fetchCategories().then(data => setCategories(data));
    }, []); // Fetch categories on mount

    const handleStart = ()=>{
        setIsRunning(true)
        startTimeRef.current = Math.floor(new Date().getTime() / 1000 )
        intervalRef.current = setInterval(()=>{
            setSecs(secs => secs + 1)
        }, 1)

    }
    const handleStop = ()=>{
        setIsRunning(false)
        clearInterval(intervalRef.current)
    }
    const handleRestart = ()=>{
        startTimeRef.current = Math.floor(new Date().getTime() / 1000 ) 
        setSecs(0)
        handleStop()
        
    }
    const handleSave = ()=>{
        console.log(secs)
        if (secs < 10*60) {
            console.log("Cannot study for less than 10 minutes")
            return
        }
        handleStop()
        
        const session = {
            start: startTimeRef.current,
            end: startTimeRef.current + secs ,
            focus: focusRef.current,
            category_id: categoryRef.current
        }
        console.log("sending session: " + JSON.stringify(session))
        postSession(session)
       
    }

    return <div className="mx-auto h-[80%] w-[80%] flex flex-col items-center justify-center gap-4">
        <div className="flex gap-5">
        {categories.length > 0 && (
            <Select onValueChange={(v)=>{categoryRef.current = v}} defaultValue={categories[0].id} >
                <SelectTrigger className="w-[180px]">
                    <SelectValue />
                </SelectTrigger>
                <SelectContent>
                    {categories.map(cat => (<SelectItem key={cat.id} value={cat.id}>{cat.name}</SelectItem>))}
                </SelectContent>
            </Select>
        )}
        <Select defaultValue="4" onValueChange={(v)=>{focusRef.current = v}}>
            <SelectTrigger className="w-[100px]">
                <SelectValue />
            </SelectTrigger>
            <SelectContent >
                {[1, 2, 3, 4, 5].map((value) => (
                <SelectItem key={value} value={String(value)}>
                    {value}
                </SelectItem>
                ))}
            </SelectContent>
        </Select>
        </div>
        <Timer secs={secs}></Timer>
        <div className="flex gap-4">
            <Button onClick={handleStart} disabled={isRunning} className="hover:bg-green-900 h-12" variant='outline'>Start</Button>
            <Button onClick={handleStop} disabled={!isRunning} className="hover:bg-red-900 h-12" variant='outline'>Stop</Button>
            <Button onClick={handleRestart} className="hover:bg-blue-950 h-12" variant='outline'>Restart</Button>
        </div>
        <div className="flex">
            <Button onClick={handleSave} className="items-center justify-center h-12">Save</Button>
        </div>

        
        
    </div>
}

async function postSession(session: Object) {
    try {
        const jsonSess = JSON.stringify(session)
        const res = await fetch("http://localhost:8080/sessions", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: jsonSess,
        })
        if (!res.ok) {
            alert("error with the sending" + res.status + " and body: " + JSON.stringify(res.body))
            return
        }
        alert("success inserting the session")
    }
    catch (e) {
        alert("caught an error: " +  e)
    }
}


