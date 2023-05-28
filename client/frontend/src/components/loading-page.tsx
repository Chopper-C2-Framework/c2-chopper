import React, { useEffect, useState } from "react";
import LoadingScreen from "../assets/loading-screen.jpg";

interface LoadingPageProps {}

const maxDots = 6;
export const LoadingPage: React.FC<LoadingPageProps> = ({}) => {
  const [dots, setDots] = useState(".");
  useEffect(() => {
    const interval = setInterval(() => {
      const newDotsNumber = dots.length + 1;
      console.log(newDotsNumber);
      if (newDotsNumber > maxDots) {
        setDots(".");
      } else {
        setDots(dots + ".");
      }
    }, 1000);

    return () => clearInterval(interval);
  }, []);

  return (
    <>
      {" "}
      <div className="absolute l-0 r-0 t-0 b-0 h-screen w-screen -z-10 ">
        <img src={LoadingScreen} className="w-screen h-screen " />
      </div>
      <div className="w-full flex justify-center items-center h-full  h-screen">
        <h1 className="text-7xl text-white  bg-black/40 font-bold font-special">
          Please wait while we verify your identity {dots}
        </h1>
      </div>
    </>
  );
};

export default LoadingPage;
