import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@components/ui/dialog";
import { zodResolver } from "@hookform/resolvers/zod";
import { usePluginsDetails } from "@hooks/queries/plugins/plugin-detail";
import { DialogClose } from "@radix-ui/react-dialog";
import { Github, Star, ToyBrick } from "lucide-react";
import { useForm } from "react-hook-form";
import * as z from "zod";
import { Badge } from "./ui/badge";
import { Button } from "./ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "./ui/card";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "./ui/form";

import { useRunPluginMutation } from "@hooks/mutations/plugin/run-plugin";
import { Input } from "./ui/input";

interface PluginCardProps {
  plugin: string;
}

export const PluginCard: React.FC<PluginCardProps> = ({ plugin }) => {
  return (
    <Card className="col-span-3">
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium font-black text-lg text-primary">
          {plugin}
        </CardTitle>
        <ToyBrick className="w-6 h-6 text-primary" />
      </CardHeader>
    </Card>
  );
};

export const LoadedPluginCard: React.FC<PluginCardProps> = ({ plugin }) => {
  const { data, error } = usePluginsDetails(plugin);
  return (
    <Card className="col-span-3">
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium font-black text-lg text-primary">
          {data?.info.Name}
        </CardTitle>
        <div className="space-x-2 flex">
          <ToyBrick className="w-6 h-6 text-primary" />
          <a
            target="_blank"
            href={data?.Metadata.source_link}
            className="cursor-pointer hover:opacity-80"
          >
            <Github className="w-6 h-6 text-secondary" />
          </a>
        </div>
      </CardHeader>
      <CardDescription className="px-8 text-left space-y-3">
        <PluginTags tags={data?.Metadata.tags || []} />
        <p> {data?.Metadata.description}</p>
      </CardDescription>
      <div className="h-4"></div>

      <CardContent>
        <InfoLabel label="Author" info={data?.Metadata.author} />
        <InfoLabel label="Version" info={data?.Metadata.version} />
        {/* <InfoLabel label="Release Date" info={data?.Metadata.release_date} /> */}
        <InfoLabel
          label="Type"
          info={mapTypeNumberToName(data?.Metadata.type)}
        />
      </CardContent>
      <CardFooter className="px-8 flex flex-col w-full ">
        <Dialog>
          <DialogTrigger className="w-full">
            <Button className="w-full " variant="outline">
              <Star className="mr-4" /> <p>Use</p>
            </Button>
          </DialogTrigger>
          <RunPluginDialog plugin={plugin} />
        </Dialog>
      </CardFooter>
    </Card>
  );
};

interface InfoLabelProps {
  label?: string;
  info?: string;
}

const InfoLabel: React.FC<InfoLabelProps> = ({ info, label }) => (
  <div className="flex items-center space-x-8 px-2">
    <p className="font-bold text-primary/90 text-md">{label}</p>
    <p>{info}</p>
  </div>
);

interface PluginTagsProps {
  tags: string[];
}

const PluginTags: React.FC<PluginTagsProps> = ({ tags }) => {
  return (
    <div className="flex items-center space-x-5 wrap">
      {tags.map((t) => (
        <Badge key={t} variant="outline">
          {t}
        </Badge>
      ))}
    </div>
  );
};

const mapTypeNumberToName = (type?: number) => {
  switch (type) {
    case 0:
      return "Info Retriever";
    case 1:
      return "Session Opener";
  }
};

const RunPluginDialog: React.FC<{ plugin: string }> = ({ plugin }) => {
  const { data } = usePluginsDetails(plugin);
  const { mutate } = useRunPluginMutation();

  const onSubmit = (data: any) => {
    mutate({
      items: data,
      file_name: plugin,
    });
  };
  return (
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Run {data?.info.Name}</DialogTitle>
        <DialogDescription></DialogDescription>
      </DialogHeader>
      <div>
        <DynamicForm
          options={data ? data.info.Options : { x: "number" }}
          onSubmit={mutate}
        />
      </div>
    </DialogContent>
  );
};

interface Options {
  [key: string]: "string" | "number";
}

const mapOptionsToSchema = (fields: Options) => {
  let schema = z.object({});
  Object.keys(fields).forEach((key) => {
    switch (key) {
      case "string":
        schema = schema.extend({
          [key]: z.string(),
        });
        break;
      case "number":
        schema = schema.extend({
          [key]: z.number(),
        });
        break;
      default:
        schema = schema.extend({
          [key]: z.string(),
        });
    }
  });

  return schema;
};

const MapOptionsToForm = ({
  fields,
  onSubmit,
}: {
  fields: Options;
  onSubmit: any;
}) => {
  const formSchema = mapOptionsToSchema(fields);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });

  // console.log(form, fields, Object.keys(fields));

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        {Object.keys(fields).map((key) => {
          // console.log("heeere", fields[key], key);
          switch (fields[key]) {
            case "string":
              return (
                <FormField
                  key={key}
                  control={form.control}
                  name={key as never}
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>{key}</FormLabel>
                      <FormControl>
                        <Input
                          placeholder={key}
                          type={fields[key] === "number" ? "number" : "string"}
                          {...field}
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              );
            case "number":
              return (
                <FormField
                  control={form.control}
                  name={key as never}
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>{key}</FormLabel>
                      <FormControl>
                        <Input
                          placeholder={key}
                          type={fields[key]}
                          {...field}
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              );
          }
        })}
        <div className="flex space-x-8 w-full items-center justify-center">
          <DialogClose>
            <Button className="min-w-[200px]" variant="outline">
              Cancel
            </Button>
          </DialogClose>
          <Button variant="secondary" className="min-w-[200px]">
            Run
          </Button>
        </div>
      </form>
    </Form>
  );
};

interface DynamicFormProps {
  options: Options;
  cancel?: () => void;
  onSubmit(data: any): void;
}

const DynamicForm: React.FC<DynamicFormProps> = ({ options, onSubmit }) => {
  function handleSubmit(data: any) {
    onSubmit(data);
  }

  return (
    <div>
      <MapOptionsToForm fields={options} onSubmit={handleSubmit} />
    </div>
  );
};

export default DynamicForm;
