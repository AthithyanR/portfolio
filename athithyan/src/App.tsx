import './App.css'
import { athi_img_b64, chennai_gmaps } from './assets'

function App() {
    return (
        <>
            <main>
                <section>
                    <article>
                        <figure>
                            <img
                                src={athi_img_b64}
                                alt='A picture of athithyan'
                            />
                            <figcaption>
                                <p>Hey, I'm Athithyan <span role='img' aria-label='wave'>👋</span></p>
                                <p className='role-caption'>
                                    <span>
                                        Senior Full Stack Developer based in&nbsp;
                                    </span>
                                    <span>
                                        <a href={chennai_gmaps} target='_blank'>
                                            Chennai, IN
                                        </a>
                                    </span>
                                </p>
                            </figcaption>
                        </figure>
                    </article>
                    <article>
                        <p>Familiar technologies:</p>
                        <ul>
                            <li>Next.js / React</li>
                            <li>Node.js</li>
                            <li>Sql / No-Sql databases</li>
                            <li>AWS</li>
                        </ul>
                        <p>Familiar languages:</p>
                        <ul>
                            <li>Typescript</li>
                            <li>Python</li>
                            <li>Golang</li>
                        </ul>
                        <p>Resume: <a href='/resume.pdf' target='_blank'>Click to view</a></p>
                    </article>
                </section>
            </main>
            <footer>
                <a href='https://github.com/AthithyanR' target='_blank' aria-label='GitHub'>
                    <svg
                        xmlns='http://www.w3.org/2000/svg'
                        width='24'
                        height='24'
                        viewBox='0 0 24 24'
                        fill='none'
                        stroke='currentColor'
                        strokeWidth='2'
                        strokeLinecap='round'
                        strokeLinejoin='round'
                        className='feather feather-github'
                    >
                        <path d='M12 2C6.48 2 2 6.48 2 12c0 4.42 2.87 8.16 6.84 9.49.5.09.66-.22.66-.49v-1.71c-2.78.6-3.37-1.34-3.37-1.34-.45-1.16-1.1-1.47-1.1-1.47-.9-.61.07-.6.07-.6 1 .07 1.53 1.03 1.53 1.03.89 1.52 2.34 1.08 2.91.83.09-.64.35-1.08.63-1.33-2.22-.25-4.56-1.11-4.56-4.92 0-1.09.39-1.99 1.03-2.7-.1-.26-.45-1.28.1-2.66 0 0 .84-.27 2.75 1.03a9.49 9.49 0 0 1 5 0c1.9-1.3 2.75-1.03 2.75-1.03.55 1.38.2 2.4.1 2.66.64.71 1.03 1.61 1.03 2.7 0 3.82-2.35 4.66-4.58 4.91.36.31.68.92.68 1.86v2.76c0 .27.17.59.68.49A10 10 0 0 0 22 12c0-5.52-4.48-10-10-10z'></path>
                    </svg>
                </a>
                <a href='https://www.linkedin.com/in/athithyan-r-008991118/' target='_blank' aria-label='LinkedIn'>
                    <svg
                        xmlns='http://www.w3.org/2000/svg'
                        width='24'
                        height='24'
                        viewBox='0 0 24 24'
                        fill='none'
                        stroke='currentColor'
                        strokeWidth='2'
                        strokeLinecap='round'
                        strokeLinejoin='round'
                        className='feather feather-linkedin'
                    >
                        <path d='M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2 2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6z'></path>
                        <rect x='2' y='9' width='4' height='12'></rect>
                        <circle cx='4' cy='4' r='2'></circle>
                    </svg>
                </a>
                <a href='mailto:athithyanr4@gmail.com' aria-label='Email'>
                    <svg
                        xmlns='http://www.w3.org/2000/svg'
                        width='24'
                        height='24'
                        viewBox='0 0 24 24'
                        fill='none'
                        stroke='currentColor'
                        strokeWidth='2'
                        strokeLinecap='round'
                        strokeLinejoin='round'
                        className='feather feather-mail'
                    >
                        <path d='M4 4h16v16H4z'></path>
                        <polyline points='22,6 12,13 2,6'></polyline>
                    </svg>
                </a>
            </footer>
        </>
    )
}

export default App
