import { GetServerSideProps } from 'next';
import React from 'react'

export default function Home() {
  return <>
    Home
  </>
}


export const getServerSideProps: any = async ({ req }) => {
  const { token } = req.cookies;

  if (!token) {
    return {
      redirect: {
        destination: '/login',
        permanent: false,
      }
    }
  }

  return {
    props: {},
  }
}