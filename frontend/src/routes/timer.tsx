import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/timer')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/timer"!</div>
}
