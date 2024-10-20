import './App.css'
import { athi_img_b64 } from './assets'

function App() {
    return (
        <main>
            <section>
                <article>
                    <figure>
                        <img
                            src={athi_img_b64}
                            alt="Athi's pic"
                        />
                        <figcaption>Hello I'm Athithyan, and this is my website.</figcaption>
                    </figure>
                    <p>I'm a Senior Full stack developer</p>
                    <p>familiar with technologies:</p>
                    <ul>
                        <li>Next.js / React</li>
                        <li>Node.js</li>
                        <li>Sql / No-Sql databases</li>
                        <li>AWS</li>
                    </ul>
                    <p>familiar with languages:</p>
                    <ul>
                        <li>Typescript</li>
                        <li>Python</li>
                        <li>Golang</li>
                    </ul>
                    {/* <br /> */}
                    <p>Resume: <a href="/resume.pdf" target="_blank">Click to view</a></p>
                    {/* <br /> */}
                    <p>Connect with me via:</p>
                    <ul>
                        <li>Github - <a href="https://github.com/AthithyanR" target="_blank">Link</a></li>
                        <li>
                            LinkedIn - <a href="https://www.linkedin.com/in/athithyan-r-008991118/" target="_blank">
                                Link
                            </a>
                        </li>
                        <li>Email - athithyanr4@gmail.com</li>
                    </ul>
                    <br />
                </article>
            </section>
        </main>
    )
}

export default App
