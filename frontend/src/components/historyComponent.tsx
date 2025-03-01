import { fetchCategories } from "@/common";
import { Category } from "@/types/types";

import { useState, useEffect } from "react";
import { RadioGroup, RadioGroupItem } from "./ui/radio-group";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "./ui/select";
import { Label } from "./ui/label";

export function HistoryComponent(){
    const [category, setCategory] = useState(1)
    const [categories, setCategories] = useState<Category[]>([]); // State for categories
    
        useEffect(() => {
            fetchCategories().then(data => setCategories(data));
        }, []); // Fetch categories on mount
        
    return <div className="flex flex-col items-center">
        <RadioGroup defaultValue="pie" className="py-5 flex gap-20">
            <div className="flex items-center space-x-2">
                <RadioGroupItem value="pie" id="pie" />
                <Label htmlFor="pie">Pie </Label>
            </div>
            <div className="flex items-center space-x-2">
                <RadioGroupItem value="graph" id="graph" />
                <Label htmlFor="graph">Graph</Label>
            </div>
        </RadioGroup>
        <div className="flex gap-10">
            <Select onValueChange={(v)=>{setCategory(Number(v))}} defaultValue={categories[0]?.id} >
                <SelectTrigger className="w-[180px]">
                    <SelectValue placeholder="Category" />
                </SelectTrigger>
                <SelectContent>
                    {categories.map(cat => (<SelectItem key={cat.id} value={cat.id}>{cat.name}</SelectItem>))}
                </SelectContent>
            </Select>
            {/* <Select onValueChange={(h)=>{setHorizon((h))}} defaultValue="current" >
                <SelectTrigger className="w-[180px]">
                    <SelectValue />
                </SelectTrigger>
                <SelectContent>
                    {timeHorizons.map((th, i) => (<SelectItem key={i} value={th}>{cat.name}</SelectItem>))}
                </SelectContent>
            </Select> */}
            <Select onValueChange={(v)=>{setCategory(Number(v))}} defaultValue={categories[0]?.id} >
                <SelectTrigger className="w-[180px]">
                    <SelectValue placeholder="Category" />
                </SelectTrigger>
                <SelectContent>
                    {categories.map(cat => (<SelectItem key={cat.id} value={cat.id}>{cat.name}</SelectItem>))}
                </SelectContent>
            </Select>
        </div>
        
        
    </div>
}

