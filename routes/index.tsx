import { Head } from "$fresh/runtime.ts";

export default function Home() {
  return (
    <>
      <Head>
        <title>Athi Portfolio</title>
      </Head>
      <div class="p-4">
        <p>I'm Athithyan, and this is my webpage.</p>
        <img
          src="/athi.webp"
          alt="athi's image"
          loading="lazy"
        />
        <p>I'm a full stack web developer mainly focusing on backend + devops.</p>
        <p>Technologies that I'm familiar with: <br></br> - React/Nextjs, Node/Deno, Sql/NoSql databases, AWS and linux.</p>
        <p>Resume: <a href="/resume.pdf" target="_blank" class="text-blue-600 visited:text-purple-600">Click to view</a></p>
        <p>Connect with me via:</p>
        <ul>
            <li>Github - <a href="https://github.com/AthithyanR" target="_blank" class="text-blue-600 visited:text-purple-600">Link</a></li>
            <li>LinkedIn - <a href="https://www.linkedin.com/in/athithyan-r-008991118/" target="_blank" class="text-blue-600 visited:text-purple-600">Link</a></li>
            <li>Email - athithyanr4@gmail.com</li>
        </ul>
      </div>
    </>
  );
}
