"use client"

import { ColumnDef } from "@tanstack/react-table"

import { Badge } from "@/components/ui/badge"
import { Checkbox } from "@/components/ui/checkbox"

import { labels, priorities, statuses } from "../data/data"
import { Task } from "../data/schema"
import { DataTableColumnHeader } from "./data-table-column-header"
import { DataTableRowActions } from "./data-table-row-actions"

export const columns: ColumnDef<Task>[] = [
  {
    id: "select",
    header: ({ table }) => (
      <Checkbox
        checked={
          table.getIsAllPageRowsSelected() ||
          (table.getIsSomePageRowsSelected() && "indeterminate")
        }
        onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
        aria-label="Select all"
        className="translate-y-[2px]"
      />
    ),
    cell: ({ row }) => (
      <Checkbox
        checked={row.getIsSelected()}
        onCheckedChange={(value) => row.toggleSelected(!!value)}
        aria-label="Select row"
        className="translate-y-[2px]"
      />
    ),
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "portMappingDescription",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Service" />
    ),
    cell: ({ row }) => <div className="w-[100px]">{row.getValue("portMappingDescription")}</div>,
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "internalClient",
    header: ({ column }) => (
        <DataTableColumnHeader column={column} title="Service IPV4" />
    ),
    cell: ({ row }) => <div className="w-[80px]">{row.getValue("internalClient")}</div>,
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "enabled",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Enabled" />
    ),
    cell: ({ row }) => {
      const status = statuses.find(
        (status) => status.value === row.getValue("enabled")
      )

      if (!status) {
        return null
      }

      return (
        <div className="flex w-[100px] items-center">
          {status.icon && (
            <status.icon className="mr-2 h-4 w-4 text-muted-foreground" />
          )}
          <span>{status.label}</span>
        </div>
      )
    },
    filterFn: (row, id, value) => {
      return value.includes(row.getValue(id))
    },
  },
  {
    accessorKey: "protocol",
    header: ({ column }) => (
        <DataTableColumnHeader column={column} title="Protocol" />
    ),
    cell: ({ row }) => <div className="w-[80px]">{row.getValue("protocol")}</div>,
    enableHiding: true,
    filterFn: (row, id, value) => {
      return value.includes(row.getValue(id))
    },
  },
  {
    accessorKey: "externalPort",
    header: ({ column }) => (
        <DataTableColumnHeader column={column} title="External port" />
    ),
    cell: ({ row }) => <div className="w-[80px]">{row.getValue("externalPort")}</div>,
    enableHiding: true,
  },
  {
    accessorKey: "internalPort",
    header: ({ column }) => (
        <DataTableColumnHeader column={column} title="Internal port" />
    ),
    cell: ({ row }) => <div className="w-[80px]">{row.getValue("internalPort")}</div>,
    enableHiding: true,
  },
  {
    id: "actions",
    cell: ({ row }) => <DataTableRowActions row={row} />,
  },
]
