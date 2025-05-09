import { main } from "motion/react-client";
import Image from "next/image";
import HeroSection from "./components/HeroSection";
import { AnimatedGridPattern } from "@/components/magicui/animated-grid-pattern";

export default function Home() {
  return (
    <main className="min-h-screen relative bg-off-white w-full ">
      <AnimatedGridPattern/>
      <HeroSection/>
    </main>
  );
}
