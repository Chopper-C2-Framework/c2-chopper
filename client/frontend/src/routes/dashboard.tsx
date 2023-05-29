import { CalendarDateRangePicker } from "@components/date-range-picker";
import {
  Activity,
  CreditCard,
  DollarSign,
  Download,
  Users,
} from "lucide-react";

import { TasksTable } from "@components/recent-sales";
import { Button } from "@components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@components/ui/card";
import { useAllPluginsQuery } from "@hooks/queries/plugins/all-plugins";
import { useAllActiveTasks } from "@hooks/queries/tasks/all-active-tasks";
import { useAllAgentsQuery } from "@hooks/queries/agents/all-agents-count";
import { useLatestTasksResults } from "@hooks/queries/tasks/get-latest-tasks-results";
import { useAllTasks } from "@hooks/queries/tasks/all-tasks";
import { Link } from "react-router-dom";
const Dashboard = () => {
  const { data: allPluginsData, isLoading: allPluginsLoading } =
    useAllPluginsQuery();
  const { data: allAgentsData, isLoading: allAgentsLoading } =
    useAllAgentsQuery();
  const { data: activeTasksData, isLoading: activeTasksLoading } =
    useAllActiveTasks();
  const { data: latestTasksData, isLoading: latestTasksLoading } =
    useLatestTasksResults({
      limit: 10,
      page: 1,
      unseen: true,
    });

  return (
    <div className="px-10">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <div className="flex items-center justify-between space-y-2">
          <h2 className="text-3xl font-bold tracking-tight">Dashboard</h2>
          <div className="flex items-center space-x-2">
            <CalendarDateRangePicker />
            <Button size="sm">
              <Download className="mr-2 h-4 w-4" />
              Download
            </Button>
          </div>
        </div>
      </div>
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4 ">
        <Link to="/app/plugins">
          <Card>
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <CardTitle className="text-sm font-medium">
                Plugins available
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">
                {allPluginsLoading || !allPluginsData
                  ? 0
                  : allPluginsData.length}
              </div>
            </CardContent>
          </Card>
        </Link>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Number of agents
            </CardTitle>
            <Users className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{allAgentsData?.count}</div>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Active tasks</CardTitle>
            <CreditCard className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{activeTasksData?.count}</div>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Unseen tasks results
            </CardTitle>
            <Activity className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{latestTasksData?.count}</div>
          </CardContent>
        </Card>
      </div>
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-7 mt-10">
        <Card className="col-span-12">
          <CardHeader>
            <CardTitle>Tasks List</CardTitle>
            <CardDescription></CardDescription>
          </CardHeader>
          <CardContent>
            <TasksTable />
          </CardContent>
        </Card>
      </div>
    </div>
  );
};

export default Dashboard;
