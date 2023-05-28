import {  ShieldAlert,LogOut, PlusCircle, Settings, User } from "lucide-react";

import { Avatar, AvatarFallback, AvatarImage } from "@components/ui/avatar";
import { Button } from "@components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuShortcut,
  DropdownMenuTrigger,
} from "@components/ui/dropdown-menu";
import { useMeQuery } from "@hooks/queries/user/me";
import { useLogout } from "@hooks/mutations/auth/useLogout";

export function UserNav() {

  const { data, isLoading } = useMeQuery()
  const mockData = {
    username: "shadcn",
    adminView: true,
  }

  const { logout }=useLogout()

  console.log(data)
   const username=!isLoading && data ?  data.user.username:mockData.username
   const isAdmin=!isLoading && data ?  data.user.role === "Admin":mockData.adminView
   const role= !isLoading && data ?  data.user.role:""

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="ghost" className="relative h-8 w-8 rounded-full">
          <Avatar className="h-8 w-8">
            <AvatarImage src="/avatars/01.png" alt="@c2-chopper" />
            <AvatarFallback>{username}</AvatarFallback>
          </Avatar>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56" align="end" forceMount>
        <DropdownMenuLabel className="font-normal">
          <div className="flex flex-col space-y-1">
            <p className="text-sm font-medium leading-none">{username}</p>
            <p className="text-xs leading-none text-muted-foreground">
              {role}
            </p>
          </div>
        </DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuGroup>
          <DropdownMenuItem>
            <User className="mr-2 h-4 w-4" />
            <span>Profile</span>
            <DropdownMenuShortcut>⇧⌘P</DropdownMenuShortcut>
          </DropdownMenuItem>
          <DropdownMenuItem>
            <ShieldAlert className="mr-2 h-4 w-4" />
            <span>Findings</span>
            <DropdownMenuShortcut>⌘B</DropdownMenuShortcut>
          </DropdownMenuItem>
          <DropdownMenuItem>
            <Settings className="mr-2 h-4 w-4" />
            <span>Settings</span>
            <DropdownMenuShortcut>⌘S</DropdownMenuShortcut>
          </DropdownMenuItem>
          {isAdmin&&
            <DropdownMenuItem>
              <PlusCircle className="mr-2 h-4 w-4" />
              <span>New Team</span>
            </DropdownMenuItem>
          }
        </DropdownMenuGroup>
        <DropdownMenuSeparator />
        <DropdownMenuItem onClick={logout}>
          <LogOut className="mr-2 h-4 w-4" />
          <span>Log out</span>
          <DropdownMenuShortcut>⇧⌘Q</DropdownMenuShortcut>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
