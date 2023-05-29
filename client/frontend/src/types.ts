import { Icons } from "@components/icons";

export type NavItem = {
  title: string;
  href: string;
  disabled?: boolean;
};

export type MainNavItem = NavItem;

export type SidebarNavItem = {
  title: string;
  disabled?: boolean;
  external?: boolean;
  icon?: keyof typeof Icons;
} & (
  | {
      href: string;
      items?: never;
    }
  | {
      href?: string;
      items: any[];
    }
);

export type SiteConfig = {
  name: string;
  description: string;
  url: string;
  ogImage: string;
  links: {
    twitter: string;
    github: string;
  };
};

export type DocsConfig = {
  mainNav: MainNavItem[];
  sidebarNav: SidebarNavItem[];
};

export type MarketingConfig = {
  mainNav: MainNavItem[];
};

export type DashboardConfig = {
  mainNav: MainNavItem[];
  sidebarNav: SidebarNavItem[];
};

export type SubscriptionPlan = {
  name: string;
  description: string;
  stripePriceId: string;
};

// export type Agent = {
//   id: string;
//   nickname: string;
//   username: string;
//   user_id: string;
//   sleep_time: number;
//   cwd: string;
// }

// export type Team = {
//   id: string;
//   name: string;
//   members: User[]
// }

// export type User = {
//   id: string;
//   username: string;
//   role: string
// }

// export enum TaskType {
//   UNKNOWN = "UNKNOWN",
//   PING = "PING",
//   SHELL = "SHELL",
// }

// export type Plugin = {
// }

export interface Agent {
  id: string;
  nickname: string;
  hostname: string;
  username: string;
  user_id: string;
  sleep_time: number;
  cwd: string;
}

export interface Team {
  id: string;
  name: string;
  members: User[];
}

export interface Plugin {
  Metadata: PluginMetadata;
  info: PluginInfo;
}

export interface PluginResult {
  id: string;
  path: string;
  output: string;
  output_type: string;
  created_at: string;
}

export interface PluginMetadata {
  version: string;
  author: string;
  tags: string[];
  release_date: string;
  type: number;
  source_link: string;
  description: string;
}

export interface PluginInfo {
  Name: string;
  Options: { [key: string]: string };
  ReturnType: string;
}

export interface Cred {
  username: string;
  password: string;
}

export enum TaskType {
  UNKNOWN = "UNKNOWN",
  PING = "PING",
  SHELL = "SHELL",
}

export interface Task {
  taskId: string;
  name: string;
  type: TaskType;
  args: string[];
  agentId: string;
  creatorId: string;
  createdAt: string;
}

export interface TaskResult {
  id: string;
  status: number;
  taskId: string;
  output: string;
  seen: boolean;
  createdAt: string;
}

export interface Host {
  hostname: string;
  ip_address: string;
  users: string[];
  creds: Cred[];
  tasks: Task[];
  used_plugins: string[];
  note: string;
}

export interface User {
  id: string;
  username: string;
  role: string;
}

export interface TaskResult {
  // ExecutedAt
  id: string;
  status: number;
  task_id: string;
  output: string;
  seen: boolean;
}
