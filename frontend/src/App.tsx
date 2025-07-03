import { Route, Routes } from "react-router-dom";

import { Register } from "./pages/Register";

function App() {
	return (
		<Routes>
			<Route element={<Register />} path="/" />
			{/*<Route element={<DocsPage />} path="/docs" />
      <Route element={<PricingPage />} path="/pricing" />
      <Route element={<BlogPage />} path="/blog" />
      <Route element={<AboutPage />} path="/about" /> */}
		</Routes>
	);
}

export default App;
