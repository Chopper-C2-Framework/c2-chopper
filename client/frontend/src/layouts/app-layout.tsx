import React from "react";
import { Icons } from "@components/icons";
import TeamSwitcher from "@components/team-switcher";
import { UserMainNav } from "@components/user-main-nav";
import { UserNav } from "@components/user-nav";
import { Button } from "@components/ui/button";
import { Search, Download } from "lucide-react";
import { Outlet, Link } from "react-router-dom";

interface AppLayoutProps {}

export const AppLayout: React.FC<AppLayoutProps> = ({}) => {
  return (
    <>
      <div className="border-b">
        <div className="flex h-16 items-center px-4">
          <Link
            to="/"
            className="flex items-center justify-center space-x-4 px-10"
          >
            <Icons.logo className="h-10 w-10" />
            <p className="font-special font-bold text-primary">C2-Chopper</p>
          </Link>
          <TeamSwitcher />
          <UserMainNav className="mx-6" />
          <div className="ml-auto flex items-center space-x-4">
            <Search />
            <UserNav />
          </div>
        </div>
      </div>{" "}
     
        {/* <Tabs defaultValue="overview" className="space-y-4">
          <TabsList>
            <TabsTrigger value="overview">Overview</TabsTrigger>
            <TabsTrigger value="analytics" disabled>
              Analytics
            </TabsTrigger>
            <TabsTrigger value="reports" disabled>
              Reports
            </TabsTrigger>
            <TabsTrigger value="notifications" disabled>
              Notifications
            </TabsTrigger>
          </TabsList>
          <TabsContent value="overview" className="space-y-4">
           
          </TabsContent>
        </Tabs> */}
        <Outlet />
    </>
  );
};

export default AppLayout;
