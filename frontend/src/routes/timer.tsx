import { createFileRoute } from '@tanstack/react-router'

import { TimerComponent } from '@/components/timerComponent'
export const Route = createFileRoute('/timer')({
  component: TimerComponent,
})


