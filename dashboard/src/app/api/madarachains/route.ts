import { NextRequest, NextResponse } from 'next/server';
import * as k8s from '@kubernetes/client-node';

const kc = new k8s.KubeConfig();
kc.loadFromDefault();

const customApi = kc.makeApiClient(k8s.CustomObjectsApi);

const GROUP = 'batch.starknet.l3';
const VERSION = 'v1alpha1';
const PLURAL = 'madarachains';
const NAMESPACE = 'default';

export async function GET() {
  try {
    const res = await customApi.listNamespacedCustomObject({
      group: GROUP,
      version: VERSION,
      namespace: NAMESPACE,
      plural: PLURAL,
    });
    console.log('GET res keys:', Object.keys(res));
    // @ts-ignore
    console.log('GET res.body:', res.body);
    // @ts-ignore
    return NextResponse.json(res.items || res.body?.items || []);
  } catch (err: any) {
    console.error('Error fetching MadaraChains:', err);
    return NextResponse.json({ error: err.message || 'Failed to fetch chains' }, { status: 500 });
  }
}

export async function POST(req: NextRequest) {
  try {
    const body = await req.json();
    const { name, replicas, chainID } = body;

    if (!name || !chainID) {
      return NextResponse.json({ error: 'Name and ChainID are required' }, { status: 400 });
    }

    const newChain = {
      apiVersion: `${GROUP}/${VERSION}`,
      kind: 'MadaraChain',
      metadata: {
        name: name,
        namespace: NAMESPACE,
      },
      spec: {
        chainID: chainID,
        replicas: replicas || 1,
        port: 9944, // Default port
      },
    };

    const res = await customApi.createNamespacedCustomObject({
      group: GROUP,
      version: VERSION,
      namespace: NAMESPACE,
      plural: PLURAL,
      body: newChain,
    });

    // @ts-ignore
    return NextResponse.json(res.body || res);
  } catch (err: any) {
    console.error('Error creating MadaraChain:', err);
    return NextResponse.json({ error: err.message || 'Failed to create chain' }, { status: 500 });
  }
}

export async function DELETE(req: NextRequest) {
  try {
    const { searchParams } = new URL(req.url);
    const name = searchParams.get('name');

    if (!name) {
      return NextResponse.json({ error: 'Name is required' }, { status: 400 });
    }

    const res = await customApi.deleteNamespacedCustomObject({
      group: GROUP,
      version: VERSION,
      namespace: NAMESPACE,
      plural: PLURAL,
      name: name,
    });

    // @ts-ignore
    return NextResponse.json(res.body || res);
  } catch (err: any) {
    console.error('Error deleting MadaraChain:', err);
    return NextResponse.json({ error: err.message || 'Failed to delete chain' }, { status: 500 });
  }
}
