import { BiLogoGmail } from "react-icons/bi"
import { FaSquareGithub } from "react-icons/fa6"

export function AboutComponent() {
    const features = [
        "1. Ability to insert a session (of study, work, whatever), specifying the focus level (from 1 to 5), and the category this session belongs to (study, work, etc)",
        `2. Ability to see your history, either in graph or pie mode, where you can see how much you have studied in specific time periods (year, month, week, day) 
            and specifying which one (current, previous, or a specific date), and for which category (study, work, backend, etc). For example, it allows getting how much one studied in the year 2023 for the "backend" category.
            in graph mode, it shows the subtotals per time frame. for example, with timemode = year, it would show something like: January 300hours, February 412 hours, and so on. 
            Whereas for the pie mode, it would show something like Backend 10000 hours, Frontend 3000 hours, Other 100 hours`,
        `3. Ability to manually insert sessions with a csv file`
    ]
    return <div className="my-2 mx-20">
        
        <p className="text-2xl"> TimeX is a straightforward productivity app to better keep track of your performance.</p>
        <p className="text-2xl mt-6"> Its core features currently are: </p>
        <ul className="ml-8 text-md">
            {features.map((f,idx) => (<li key={idx}>{f}</li>))}
        </ul>
        <p className="text-2xl mt-6"> Its Tech Stack: </p>
        <ul className="ml-8 text-md">
            <li>Backend: Go v1.24 & Chi</li>
            <li>Frontend: React v19, Typescript, Tailwindcss, Tanstack Router & Vite</li>
            <li>Database: PostgreSQL v17</li>
            <li>Version Control: Git</li>
            <li>Containerization: Docker</li>
        </ul>
        <p className="text-2xl mt-6">
            About the creator: 
        </p>
        <p className="ml-8 text-lg">
            Hi there! I'm Erik Rodriguez Novikov, a Software Developer with emphasis on backend-dev (with golang & typescript), 
            with a passion for problem-solving and software-dev in general.
        </p>
        <p className="text-2xl mt-2">Socials: </p>
        <p className="ml-8"><a href="https://github.com/E-nkv">Github <FaSquareGithub className="inline-block" size={30} /> github.com/E-nkv </a></p>
        <p className="ml-8"><a href="mailto:erikrn24@gmail.com"> Gmail <BiLogoGmail size={30} className="inline-block"/> erikrn24@gmail.com</a> </p>

        
    </div>
}