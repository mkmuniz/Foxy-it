import React from 'react'
import HomeContent from '../components/homeContent';
import Navbar from '../components/navbar';
import withAuth from '../utils/auth/withAuth'

function Home() {
  return <>
      <Navbar />
      <HomeContent />
  </>
}

export default withAuth(Home);