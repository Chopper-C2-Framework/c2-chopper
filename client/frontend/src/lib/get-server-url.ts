export const getServerUrl = ():string => {
    console.log(import.meta.env)
    return import.meta.env.DEV ? "http://localhost:4000" : import.meta.env.VITE_API_URL as string;
}