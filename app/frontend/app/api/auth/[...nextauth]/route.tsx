import NextAuth from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import { hash, compare } from "bcryptjs";

const getAuthenticationByEmail = async (email: string) => {
  const apiUrl = `${process.env.API_URL}/authentication/search?email=${email}`;
  const res = await fetch(apiUrl);

  if (!res.ok) {
    throw new Error("Failed to fetch authentication info");
  }
  const data = await res.json();

  return {
    id: String(data.ID),
    email: data.Email,
    password: data.Password,
  };
};

export const authOptions = {
  providers: [
    CredentialsProvider({
      name: "Credentials",
      credentials: {
        email: { label: "Email", type: "text" },
        password: { label: "Password", type: "password" },
      },
      async authorize(credentials) {
        console.log("aaaaa");
        console.log("credentials!.email", credentials!.email);
        const authentication = await getAuthenticationByEmail(
          credentials!.email
        );
        if (!authentication) {
          return null;
        }

        const isValid = await compare(
          credentials!.password,
          authentication.password
        );
        if (!isValid) {
          return null;
        }

        return { id: authentication.id, email: authentication.email };
      },
    }),
  ],
  pages: {
    signIn: "/login",
    error: "/login",
  },
};

const handler = NextAuth(authOptions);
export { handler as GET, handler as POST };
