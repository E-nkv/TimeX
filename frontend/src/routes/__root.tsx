import { createRootRoute, Link, Outlet } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'
import '../App.css'

const links = [
  {to: "/timer", text: "Timer"},
  {to: "/history", text: "History"},
  {to: "/about", text: "About"},
]
export const Route = createRootRoute({
  component: () => (
    <div className='bg-[--background] h-svh'>
     <div className="flex gap-2 divide-x divide-gray-600">
        {links.map((link)=> (
          <Link
            key={link.to}
            to={link.to}
            className="px-20 py-2 text-xl text-gray-400 rounded-md [&.active]:text-white [&.active]:font-bold [&.active]:bg-gray-900"
          >
            {link.text}
          </Link>
        ))}
      </div>
      <hr />
      <Outlet />
      <TanStackRouterDevtools />
    </div>
  ),
})