import { z } from "zod"

// We're keeping a simple non-relational schema here.
// IRL, you will have a schema for your data models.
export const taskSchema = z.object({
  portMappingDescription: z.string(),
  internalClient: z.string(),
  enabled: z.boolean(),
  protocol: z.string(),
  externalPort: z.number(),
  internalPort: z.number(),
})

export type Task = z.infer<typeof taskSchema>
