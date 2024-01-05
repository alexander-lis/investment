import Header from "./Header";
import Footer from "./Footer";
import SubHeader from "./SubHeader";

export default function Root() {
  return (
    <div className="h-screen w-screen flex flex-col bg-gray-100">
      <Header />
      <SubHeader></SubHeader>
      <div className="flex-auto">body</div>
      <Footer />
    </div>
  );
}
