package com.bandcamp.android.network;
public class d implements java.util.Observer {
    private static String e = "X";
    private static String f = "aB";
    private static String g = "pmac";
    private static String h = "D";
    private static String i = "M";
    private static char[] j;
    private byte[] a;
    private int b;
    private byte[] c;
    private int d;

    static d()
    {
        char[] v0_1 = new char[5];
        com.bandcamp.android.network.d.j = v0_1;
        char[] v0_8 = new StringBuilder();
        v0_8.append(com.bandcamp.android.network.d.g);
        v0_8.append("d");
        com.bandcamp.android.network.d.g = v0_8.toString();
        char[] v0_4 = new StringBuilder();
        v0_4.append("n");
        v0_4.append(com.bandcamp.android.network.d.f);
        com.bandcamp.android.network.d.f = v0_4.toString();
        char[] v0_6 = com.bandcamp.android.network.d.j;
        v0_6[0] = 100;
        v0_6[1] = 116;
        v0_6[2] = 109;
        v0_6[3] = 102;
        v0_6[4] = 97;
        return;
    }

    public d(com.bandcamp.android.network.CacheListener p2)
    {
        this.b = -1;
        this.c = "".getBytes();
        this.d = 0;
        p2.getObservable().a(this);
        return;
    }

    private String a(byte[] body)
    {
        return this.a(this.a(), body);
    }

    private byte[] a()
    {
        byte[] v0_0 = this.a;
        if (v0_0 == null) {
            v0_0 = new byte[0];
        }
        return v0_0;
    }

    public String a(byte[] p4, byte[] body)
    {
        if (body != null) {
            byte[] v1_6 = new byte[(p4.length + body.length)];
            System.arraycopy(p4, 0, v1_6, 0, p4.length);
            System.arraycopy(body, 0, v1_6, p4.length, body.length);
            p4 = v1_6;
        }
        String v5_1 = this.d;
        if (v5_1 <= null) {
            return com.bandcamp.shared.util.i.b(new String(p4), new String(com.bandcamp.android.network.d.j));
        } else {
            String v4_6;
            if (v5_1 <= 30) {
                v4_6 = com.bandcamp.shared.util.i.a(new String(p4), new String(this.c), 0);
            } else {
                v4_6 = com.bandcamp.shared.util.i.a(new String(p4), new String(this.c), 0, 0);
            }
            return v4_6;
        }
    }

    public void a(com.bandcamp.shared.network.d p8)
    {
        try {
            String v2_5 = com.bandcamp.android.network.d.f.toCharArray();
            String v3_3 = 0;
        } catch (java.io.IOException v8_1) {
            Object[] v1_1 = new Object[0];
            com.bandcamp.shared.util.BCLog.c.b(v8_1, v1_1);
            return;
        } catch (java.io.IOException v8_1) {
        } catch (java.io.IOException v8_1) {
        }
        while (v3_3 < (v2_5.length / 2)) {
            StringBuilder v4_3 = v2_5[v3_3];
            v2_5[v3_3] = v2_5[((v2_5.length - v3_3) - 1)];
            v2_5[((v2_5.length - v3_3) - 1)] = v4_3;
            v3_3++;
        }
        String v3_0 = new String(v2_5);
        String v2_1 = com.bandcamp.android.network.d.g.toCharArray();
        StringBuilder v4_0 = 0;
        while (v4_0 < (v2_1.length / 2)) {
            String v5_3 = v2_1[v4_0];
            v2_1[v4_0] = v2_1[((v2_1.length - v4_0) - 1)];
            v2_1[((v2_1.length - v4_0) - 1)] = v5_3;
            v4_0++;
        }
        StringBuilder v4_2 = new StringBuilder();
        v4_2.append(com.bandcamp.android.network.d.e);
        v4_2.append("-");
        v4_2.append(v3_0);
        v4_2.append(new String(v2_1));
        v4_2.append("-");
        v4_2.append(com.bandcamp.android.network.d.h);
        v4_2.append(com.bandcamp.android.network.d.i);
        p8.b(v4_2.toString(), this.a(p8.d()));
        return;
    }

    public void update(java.util.Observable p4, Object p5)
    {
        if (!(p5 instanceof com.bandcamp.android.network.CacheListenerEvent)) {
            if (((p5 instanceof String)) && ((this.b + 48) == ((String) p5).charAt(0))) {
                try {
                    this.a = p5.toString().substring(1).getBytes("utf-8");
                } catch (java.io.UnsupportedEncodingException v4_4) {
                    Object[] v0_1 = new Object[0];
                    com.bandcamp.shared.util.BCLog.c.b(v4_4, v0_1);
                }
            }
        } else {
            java.io.UnsupportedEncodingException v4_8 = com.bandcamp.shared.util.e.a(((com.bandcamp.android.network.CacheListenerEvent) p5).value());
            if (v4_8 != -1) {
                this.b = v4_8;
                if (v4_8 != 3) {
                    if (v4_8 != 4) {
                        this.d = 0;
                    } else {
                        this.c = com.bandcamp.shared.platform.e.a().a(((com.bandcamp.android.network.CacheListenerEvent) p5).value(), 0).getBytes("utf-8");
                        this.d = 60;
                    }
                } else {
                    this.c = com.bandcamp.shared.platform.e.a().b(((com.bandcamp.android.network.CacheListenerEvent) p5).value()).getBytes("utf-8");
                    this.d = 30;
                }
            }
        }
        return;
    }
}
