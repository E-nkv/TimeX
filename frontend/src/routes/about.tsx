import { createFileRoute } from '@tanstack/react-router'
import { AboutComponent } from '@/components/aboutComponent'

export const Route = createFileRoute('/about')({
  component: AboutComponent,
})


