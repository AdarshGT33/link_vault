import { Button } from "@/components/ui/button";
import { CardStack } from "@/components/ui/card-stack";
import { FlipWords } from "@/components/ui/flip-words";
import Link from "next/link";

export default function HeroSection (){
    const words = ["beautifully!", "your way!"]
    const cards = [
        {
            id: 0,
            imageUrl:"https://images.unsplash.com/photo-1746469471553-afa9f34157fd?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxmZWF0dXJlZC1waG90b3MtZmVlZHwzfHx8ZW58MHx8fHx8",
        },
        {
            id: 1,
            imageUrl:"https://images.unsplash.com/photo-1744360820043-59b729149f7d?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxmZWF0dXJlZC1waG90b3MtZmVlZHw0M3x8fGVufDB8fHx8fA%3D%3D",
        },
        {
            id: 2,
            imageUrl:"https://images.unsplash.com/photo-1745810187217-4d9e1ccfd9d5?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxmZWF0dXJlZC1waG90b3MtZmVlZHw0N3x8fGVufDB8fHx8fA%3D%3D",
        },
    ];
    
    return (
        <>
        <div className="absolute h-8/12 w-11/12 top-8 left-10 bg-mocha rounded-b-lg rounded-tl-lg flex items-center justify-end">
            <div className="h-full w-1/2 flex items-center justify-center flex-col">
                <div className="w-full h-4/6 pt-2 flex items-center">
                    <h1 className="text-cream text-7xl">
                        Organize Your Links, <FlipWords words={words}/>
                    </h1>
                </div>
                <div className="w-full h-1/6 p-1.5">
                    <p className="font-edu text-md text-muted">
                        Share your digital footprint with style. <span className="font-cal text-off-white font-bold">LinkVault</span> keeps it clean and simple.
                    </p>
                </div>
                <div className=" h-1/6 w-full flex items-center justify-evenly pr-40 gap-0">
                    <Link href={"/signup"}>
                        <Button>
                            Sign Up
                        </Button>
                    </Link>
                    <Link href={"/video"}>
                        <Button variant={"secondary"}>
                            View Demo
                        </Button>
                    </Link>
                </div>
            </div>
            <div className="h-full w-1/2 flex items-center justify-center">
                <CardStack items={cards}/>
            </div>
        </div>
        </>
    )
}