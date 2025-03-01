import { createFileRoute } from '@tanstack/react-router'
import { HistoryComponent } from '@/components/historyComponent'

export const Route = createFileRoute('/history')({
  component: HistoryComponent
})


