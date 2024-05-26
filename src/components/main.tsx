import { Component, JSX } from "solid-js";
import Navbar from "./navbar";

interface Props {
  children: JSX.Element | JSX.Element[]
}

const Main: Component<Props> = ({ children }) => {
  return <>
    <main class="max-w-screen-2xl h-screen relative">
      <Navbar />
      {children}
    </main>
  </>
}

export default Main