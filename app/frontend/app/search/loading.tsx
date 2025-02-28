// const Loading = () => {
//   return (
//     <div className="flex justify-center items-center gap-6 mt-10">
//       <div className="h-10 w-10 animate-spin border-[5px] border-sky-400 rounded-full  border-t-transparent"></div>
//       <p className="text-[30px] font-weight">Loading</p>
//     </div>
//   );
// };

// export default Loading;

export default function Loading() {
  return (
    <div className="flex justify-center items-center h-screen">
      <div className="w-10 h-10 border-4 border-blue-500 border-t-transparent rounded-full animate-spin"></div>
      <p className="ml-4 text-lg font-semibold">Loading...</p>
    </div>
  );
}
