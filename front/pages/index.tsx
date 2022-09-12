import React from 'react'
import Footer from '../components/footer';
import HomeContent from '../components/homeContent';
import Navbar from '../components/navbar';
import withAuth from '../utils/auth/withAuth'

function Home() {
  return <>
      <Navbar />
      <HomeContent />
      <Footer />
  </>
}

export default withAuth(Home);